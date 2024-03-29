import axios from "axios";

import type { StoneMovement } from "../types/stoneMovement";
import {rootURL} from "../settings/settings"


//ユーザーのゲーム盤面を新規作成
export const postBoardData = async ()=>{
  const response =await axios.post(rootURL)
  return response.data
}

//ユーザーのゲーム盤面を取得
export const getBoardData = async (boardId:number)=>{
  const response = await axios.get(`${rootURL}${boardId}`)
  return response.data
}

//ユーザーが石を置いて盤面を更新する処理
export const updateStonePos = async (boardId:number,stoneMovement:StoneMovement)=>{
  const response = await axios.put(`${rootURL}${boardId}`,
    {is_my_turn:true,
    ...stoneMovement},
    {headers: {
      "Content-Type": "multipart/form-data",
    }})
  return response.data
}

//ユーザーの対戦相手（コンピューター）が石を置いて盤面を更新する処理
export const updateStonePosByOpponent = async (boardId:number,color:number)=>{
  const response = await axios.put(`${rootURL}${boardId}`,
    { is_my_turn:false,
      my_stone_color:color
    },
    {headers: {
      "Content-Type": "multipart/form-data",
    }})
  return response.data
}