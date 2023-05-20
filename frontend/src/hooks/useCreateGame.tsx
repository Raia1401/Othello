import React from "react"
import { useEffect,useContext,useState} from "react";
import { BoardDataContext} from "../providers/boardDataProvider";
import * as apis from "../apis/api";
import type { BoardData } from "../types/boardData";


export const useCreateGame=()=>{

  // const {boardId,setBoardId}=useContext(BoardDataContext)
  const [boardId,setBoardId]=useState(0)

  const onClick = () =>{
      apis.postBoardData().then((boardData:BoardData)=>{
        setBoardId(boardData.data.BoardId)
        
      })
  }

  return {boardId,onClick}

}
