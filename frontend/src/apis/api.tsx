import axios from "axios";
import type { StoneMovement } from "../types/StoneMovement";

const rootURL="http://localhost:3030/othello/v1/gamematch"


//対応するユーザーのゲーム盤面を新規作成
export const postBoardData = async (userId:number)=>{
  const response =await axios.post(rootURL,{"user_id":userId})
  return response.data
}

//対応するユーザーのゲーム盤面を取得
export const getBoardData = async (userId:number)=>{
  const response = await axios.get(`${rootURL}/${userId}`)
  return response.data
}

//対応するユーザーが石を置いて盤面を更新する処理
export const updateStonePos = async (matchId:number,stoneMovement:StoneMovement)=>{
  const response = await axios.put(`${rootURL}/myself/${matchId}`,stoneMovement)
  return response.data
}

//ユーザーの対戦相手（コンピューター）が石を置いて盤面を更新する処理
export const updateStonePosByOpponent = async (matchId:number)=>{
  const response = await axios.put(`${rootURL}/opponent/${matchId}`)
  return response.data
}