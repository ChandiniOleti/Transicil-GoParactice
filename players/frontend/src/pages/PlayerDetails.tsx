import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getPlayerById } from "../services/playerApi";
import type { Player } from "../types/player";
import "./PlayerDetails.css";

function PlayerDetails() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [player, setPlayer] = useState<Player | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchPlayer = async () => {
      if (!id) {
        setError("Player ID is missing.");
        setLoading(false);
        return;
      }

      try {
        setLoading(true);
        setError("");
        setPlayer(await getPlayerById(id));
      } catch {
        setError("Failed to load player details.");
      } finally {
        setLoading(false);
      }
    };

    fetchPlayer();
  }, [id]);

  if (loading) return <p>Loading player details...</p>;

  if (error || !player) {
    return (
      <div>
        <p>{error || "Player not found."}</p>
        <button
  className="back-button"
  onClick={() => navigate("/players")}
>
  ← Back to Players
</button>
      </div>
    );
  }

  const details: Array<[string, string | number | null]> = [
    ["Player ID", player.playerID],
    ["First Name", player.nameFirst],
    ["Last Name", player.nameLast],
    ["Given Name", player.nameGiven],
    ["Birth Year", player.birthYear],
    ["Birth Month", player.birthMonth],
    ["Birth Day", player.birthDay],
    ["Birth Country", player.birthCountry],
    ["Birth State", player.birthState],
    ["Birth City", player.birthCity],
    ["Death Year", player.deathYear],
    ["Death Month", player.deathMonth],
    ["Death Day", player.deathDay],
    ["Death Country", player.deathCountry],
    ["Death State", player.deathState],
    ["Death City", player.deathCity],
    ["Height", player.height ? `${player.height} in` : null],
    ["Weight", player.weight ? `${player.weight} lbs` : null],
    ["Bats", player.bats],
    ["Throws", player.throws],
    ["Debut", player.debut],
    ["Final Game", player.finalGame],
    ["Retro ID", player.retroID],
    ["BBRef ID", player.bbrefID],
  ];

  return (
    <div className="player-details-page">
      <button onClick={() => navigate("/players")}>
        ← Back to Players
      </button>

      <h1>Player Details</h1>

      <div className="details-grid">
        {details.map(([label, value]) => (
          <div className="detail-item" key={label}>
            <span>{label}</span>
            <strong>{value ?? "-"}</strong>
          </div>
        ))}
      </div>
    </div>
  );
}

export default PlayerDetails;