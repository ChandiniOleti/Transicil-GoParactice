import { useEffect, useState, type FormEvent } from "react";
import { Link, useNavigate } from "react-router-dom";
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
  const [searchPlayerId, setSearchPlayerId] = useState("");
  const [goToPage, setGoToPage] = useState("");

  const limit = 20;

  useEffect(() => {
    const fetchPlayers = async () => {
      try {
        setLoading(true);
        setError("");

        const data = await getPlayers(page, limit);

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
  }, [page]);

  const handleSearch = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const playerId = searchPlayerId.trim();

    if (playerId) {
      navigate(`/players/${encodeURIComponent(playerId)}`);
    }
  };

  const handleGoToPage = (event: FormEvent<HTMLFormElement>) => {
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

  return (
    <div className="players-page">
      <div className="players-container">
        <div className="players-header">
          <div>
            <h1>Players</h1>

            <p className="page-info">
              Browse and search player records
            </p>
          </div>

          <form className="player-search" onSubmit={handleSearch}>
            <input
              type="text"
              value={searchPlayerId}
              onChange={(event) => setSearchPlayerId(event.target.value)}
              placeholder="Enter Player ID, e.g. aardsda01"
              aria-label="Player ID"
            />

            <button type="submit">Search</button>
          </form>
        </div>

        {loading && (
          <p className="status-message">
            Loading players...
          </p>
        )}

        {error && (
          <p className="error-message">
            {error}
          </p>
        )}

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
                      <td colSpan={9} className="no-data">
                        No players found
                      </td>
                    </tr>
                  ) : (
                    players.map((player) => (
                      <tr key={player.playerID}>
                        <td className="player-id">
                          <Link
                            to={`/players/${encodeURIComponent(
                              player.playerID
                            )}`}
                          >
                            {player.playerID || "-"}
                          </Link>
                        </td>

                        <td>{player.nameFirst || "-"}</td>
                        <td>{player.nameLast || "-"}</td>
                        <td>{player.birthYear ?? "-"}</td>
                        <td>{player.birthCountry || "-"}</td>
                        <td>{player.height ?? "-"}</td>
                        <td>{player.weight ?? "-"}</td>
                        <td className="center">{player.bats || "-"}</td>
                        <td className="center">{player.throws || "-"}</td>
                      </tr>
                    ))
                  )}
                </tbody>
              </table>
            </div>

            <div className="pagination">
              <button
                onClick={() =>
                  setPage((current) => Math.max(current - 1, 1))
                }
                disabled={page === 1}
              >
                Previous
              </button>

              <span>
                Page {page} of {totalPages}
              </span>

              <form className="page-jump" onSubmit={handleGoToPage}>
                <input
                  type="number"
                  min="1"
                  max={totalPages}
                  value={goToPage}
                  onChange={(event) => setGoToPage(event.target.value)}
                  placeholder="Go to page"
                  aria-label="Go to page"
                />

                <button type="submit">Go</button>
              </form>

              <button
                onClick={() =>
                  setPage((current) =>
                    Math.min(current + 1, totalPages)
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