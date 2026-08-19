import { useEffect, useState, type FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import { getPlayers } from "../services/playerApi";
import type { Player } from "../types/player";
import "./Players.css";

function Players() {
  const navigate = useNavigate();

  const [players, setPlayers] = useState<Player[]>([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  // What user is currently typing
  const [searchName, setSearchName] = useState("");

  // Actual search value used for API
  const [activeSearch, setActiveSearch] = useState("");

  // Page number entered by user
  const [goToPage, setGoToPage] = useState("");

  const limit = 20;

  useEffect(() => {
    const fetchPlayers = async () => {
      try {
        setLoading(true);
        setError("");

        const data = await getPlayers(
          page,
          limit,
          activeSearch
        );

        setPlayers(data.players || []);
        setTotalPages(data.totalPages || 1);
      } catch (err) {
        console.error("Failed to fetch players:", err);
        setError("Failed to load players.");
        setPlayers([]);
      } finally {
        setLoading(false);
      }
    };

    fetchPlayers();
  }, [page, activeSearch]);

  // Runs when Search button is clicked
  const handleSearch = (
    event: FormEvent<HTMLFormElement>
  ) => {
    event.preventDefault();

    const name = searchName.trim();

    if (name.length < 2) {
      setError("Please enter at least 2 letters to search.");
      return;
    }

    setError("");
    setPage(1);
    setActiveSearch(name);
  };

  // Runs while user is typing
  const handleSearchChange = (value: string) => {
    setSearchName(value);
    setError("");

    const name = value.trim();

    // Empty input → show all players
    if (name.length === 0) {
      setPage(1);
      setActiveSearch("");
      return;
    }

    // As soon as 2 letters are typed → search automatically
    if (name.length >= 2) {
      setPage(1);
      setActiveSearch(name);
    }
  };

  const handleClearSearch = () => {
    setSearchName("");
    setActiveSearch("");
    setPage(1);
    setError("");
  };

  const handleGoToPage = (
    event: FormEvent<HTMLFormElement>
  ) => {
    event.preventDefault();

    const requestedPage = Number(goToPage);

    if (
      Number.isInteger(requestedPage) &&
      requestedPage >= 1 &&
      requestedPage <= totalPages
    ) {
      setPage(requestedPage);
      setGoToPage("");
    }
  };

  // Click anywhere on a player's row
  const handlePlayerClick = (playerID: string) => {
    if (!playerID) {
      return;
    }

    navigate(
      `/players/${encodeURIComponent(playerID)}`
    );
  };

  return (
    <div className="players-page">
      <div className="players-container">

        {/* Header */}
        <div className="players-header">
          <div>
            <h1>Players</h1>

            <p className="page-info">
              {activeSearch
                ? `Search results for "${activeSearch}"`
                : "Browse and search player records"}
            </p>
          </div>

          {/* Search */}
          <form
            className="player-search"
            onSubmit={handleSearch}
          >
            <input
              type="text"
              value={searchName}
              onChange={(event) =>
                handleSearchChange(event.target.value)
              }
              placeholder="Search first or last name"
              aria-label="Search player name"
            />

            <button type="submit">
              Search
            </button>

            {searchName && (
              <button
                type="button"
                onClick={handleClearSearch}
              >
                Clear
              </button>
            )}
          </form>
        </div>

        {/* Loading */}
        {loading && (
          <p className="status-message">
            Loading players...
          </p>
        )}

        {/* Error */}
        {error && (
          <p className="error-message">
            {error}
          </p>
        )}

        {/* Players */}
        {!loading && !error && (
          <>
            <div className="table-wrapper">
              <table className="players-table">

                <thead>
                  <tr>
                    <th>Player ID</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Birth Year</th>
                    <th>Birth Country</th>
                    <th>Height</th>
                    <th>Weight</th>
                    <th>Bats</th>
                    <th>Throws</th>
                  </tr>
                </thead>

                <tbody>
                  {players.length === 0 ? (
                    <tr>
                      <td
                        colSpan={9}
                        className="no-data"
                      >
                        No players found
                      </td>
                    </tr>
                  ) : (
                    players.map((player) => (
                      <tr
                        key={player.playerID}
                        className="clickable-row"
                        onClick={() =>
                          handlePlayerClick(
                            player.playerID
                          )
                        }
                      >
                        <td className="player-id">
                          {player.playerID || "-"}
                        </td>

                        <td>
                          {player.nameFirst || "-"}
                        </td>

                        <td>
                          {player.nameLast || "-"}
                        </td>

                        <td>
                          {player.birthYear ?? "-"}
                        </td>

                        <td>
                          {player.birthCountry || "-"}
                        </td>

                        <td>
                          {player.height ?? "-"}
                        </td>

                        <td>
                          {player.weight ?? "-"}
                        </td>

                        <td className="center">
                          {player.bats || "-"}
                        </td>

                        <td className="center">
                          {player.throws || "-"}
                        </td>
                      </tr>
                    ))
                  )}
                </tbody>

              </table>
            </div>

            {/* Pagination */}
            <div className="pagination">

              <button
                onClick={() =>
                  setPage((current) =>
                    Math.max(current - 1, 1)
                  )
                }
                disabled={page === 1}
              >
                Previous
              </button>

              <span>
                Page {page} of {totalPages}
              </span>

              <form
                className="page-jump"
                onSubmit={handleGoToPage}
              >
                <input
                  type="number"
                  min="1"
                  max={totalPages}
                  value={goToPage}
                  onChange={(event) =>
                    setGoToPage(event.target.value)
                  }
                  placeholder="Go to page"
                  aria-label="Go to page"
                />

                <button type="submit">
                  Go
                </button>
              </form>

              <button
                onClick={() =>
                  setPage((current) =>
                    Math.min(
                      current + 1,
                      totalPages
                    )
                  )
                }
                disabled={page === totalPages}
              >
                Next
              </button>

            </div>
          </>
        )}

      </div>
    </div>
  );
}

export default Players;