import { useEffect,useContext,useState} from "react";
import { BoardDataContext} from "../providers/boardDataProvider";
import * as apis from "../apis/api";
import type { BoardData } from "../types/boardData";


export const useNewGame=()=>{

  const boardDataCtx=useContext(BoardDataContext)

  const createGame = () =>{
      apis.postBoardData().then((boardData:BoardData)=>{
        boardDataCtx.setBoardId(boardData.data.BoardId)
      })
  }

  return createGame
}
