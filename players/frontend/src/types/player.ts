export interface Player {
  playerID: string;

  birthYear: number | null;
  birthMonth: number | null;
  birthDay: number | null;
  birthCountry: string | null;
  birthState: string | null;
  birthCity: string | null;

  deathYear: number | null;
  deathMonth: number | null;
  deathDay: number | null;
  deathCountry: string | null;
  deathState: string | null;
  deathCity: string | null;

  nameFirst: string | null;
  nameLast: string | null;
  nameGiven: string | null;

  weight: number | null;
  height: number | null;
  bats: string | null;
  throws: string | null;

  debut: string | null;
  finalGame: string | null;
  retroID: string | null;
  bbrefID: string | null;
}

export interface PlayersResponse {
  limit: number;
  offset: number;
  page: number;
  players: Player[];
  total: number;
  totalPages: number;
}