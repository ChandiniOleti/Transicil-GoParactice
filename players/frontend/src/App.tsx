import { Navigate, Route, Routes } from "react-router-dom";
import Players from "./pages/Players";
import PlayerDetails from "./pages/PlayerDetails";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/players" replace />} />
      <Route path="/players" element={<Players />} />
      <Route path="/players/:id" element={<PlayerDetails />} />
    </Routes>
  );
}

export default App;