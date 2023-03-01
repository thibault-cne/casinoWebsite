import { User } from "./user";

export type RouletteResult = {
    id: string;
    game_id: string;
    color: string;
    amount: number;
    win: boolean;
}

export type RouletteBet = {
    id: string;
    game_id: string;
    user: User;
    color: string;
    amount: number;
};