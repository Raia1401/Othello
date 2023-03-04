import axios from "axios";
import type { StoneMovement } from "../types/StoneMovement";

const rootURL:string = process.env.REACT_APP_API || ""


//ユーザーのゲーム盤面を新規作成
export const postBoardData = async ()=>{
  const response =await axios.post(rootURL)
  return response.data
}

//ユーザーのゲーム盤面を取得
export const getBoardData = async (boardId:number)=>{
  const response = await axios.get(`${rootURL}/${boardId}`)
  return response.data
}

//ユーザーが石を置いて盤面を更新する処理
export const updateStonePos = async (boardId:number,stoneMovement:StoneMovement)=>{
  console.log(stoneMovement)
  const response = await axios.put(`${rootURL}/myself/${boardId}`,
    stoneMovement,
    {headers: {
      "Content-Type": "multipart/form-data",
    }})
  return response.data
}

//ユーザーの対戦相手（コンピューター）が石を置いて盤面を更新する処理
export const updateStonePosByOpponent = async (boardId:number,color:number)=>{
  const response = await axios.put(`${rootURL}/opponent/${boardId}`,
    {color:color},
    {headers: {
      "Content-Type": "multipart/form-data",
    }})
  return response.data
}