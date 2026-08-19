import api from "./api";
import type { Player, PlayersResponse } from "../types/player";

type RawPlayer = Record<string, unknown>;

type RawPlayersResponse = Omit<PlayersResponse, "players"> & {
  players?: RawPlayer[];
};

const readNumber = (value: unknown): number | null => {
  if (typeof value === "number") {
    return value;
  }

  if (value && typeof value === "object") {
    const data = value as {
      Int32?: unknown;
      Valid?: boolean;
    };

    if (data.Valid === false) {
      return null;
    }

    if (typeof data.Int32 === "number") {
      return data.Int32;
    }
  }

  return null;
};

const readString = (value: unknown): string | null => {
  if (typeof value === "string") {
    return value;
  }

  if (value && typeof value === "object") {
    const data = value as {
      String?: unknown;
      Valid?: boolean;
    };

    if (data.Valid === false) {
      return null;
    }

    if (typeof data.String === "string") {
      return data.String;
    }
  }

  return null;
};

const readDate = (value: unknown): string | null => {
  if (typeof value === "string") {
    return value;
  }

  if (value && typeof value === "object") {
    const data = value as {
      Time?: unknown;
      Valid?: boolean;
    };

    if (data.Valid === false) {
      return null;
    }

    if (typeof data.Time === "string") {
      return data.Time;
    }
  }

  return null;
};

const mapPlayer = (player: RawPlayer): Player => ({
  playerID: readString(
    player.playerID ?? player.Playerid
  ) ?? "",

  birthYear: readNumber(
    player.birthYear ?? player.Birthyear
  ),

  birthMonth: readNumber(
    player.birthMonth ?? player.Birthmonth
  ),

  birthDay: readNumber(
    player.birthDay ?? player.Birthday
  ),

  birthCountry: readString(
    player.birthCountry ?? player.Birthcountry
  ),

  birthState: readString(
    player.birthState ?? player.Birthstate
  ),

  birthCity: readString(
    player.birthCity ?? player.Birthcity
  ),

  deathYear: readNumber(
    player.deathYear ?? player.Deathyear
  ),

  deathMonth: readNumber(
    player.deathMonth ?? player.Deathmonth
  ),

  deathDay: readNumber(
    player.deathDay ?? player.Deathday
  ),

  deathCountry: readString(
    player.deathCountry ?? player.Deathcountry
  ),

  deathState: readString(
    player.deathState ?? player.Deathstate
  ),

  deathCity: readString(
    player.deathCity ?? player.Deathcity
  ),

  nameFirst: readString(
    player.nameFirst ?? player.Namefirst
  ),

  nameLast: readString(
    player.nameLast ?? player.Namelast
  ),

  nameGiven: readString(
    player.nameGiven ?? player.Namegiven
  ),

  weight: readNumber(
    player.weight ?? player.Weight
  ),

  height: readNumber(
    player.height ?? player.Height
  ),

  bats: readString(
    player.bats ?? player.Bats
  ),

  throws: readString(
    player.throws ?? player.Throws
  ),

  debut: readDate(
    player.debut ?? player.Debut
  ),

  finalGame: readDate(
    player.finalGame ?? player.Finalgame
  ),

  retroID: readString(
    player.retroID ?? player.Retroid
  ),

  bbrefID: readString(
    player.bbrefID ?? player.Bbrefid
  ),
});

export const getPlayers = async (
  page: number,
  limit: number,
  searchName: string = ""
): Promise<PlayersResponse> => {
  const response = await api.get<RawPlayersResponse>(
    "/players",
    {
      params: {
        page,
        limit,

        // Backend expects ?name=
        name: searchName || undefined,
      },
    }
  );

  return {
    ...response.data,
    players: (response.data.players ?? []).map(mapPlayer),
  };
};

export const getPlayerById = async (
  playerId: string
): Promise<Player> => {
  const response = await api.get<RawPlayer>(
    `/players/${encodeURIComponent(playerId)}`
  );

  return mapPlayer(response.data);
};