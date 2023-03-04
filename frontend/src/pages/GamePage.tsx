import Grid from "../componetns/Grid";
import { useBoardData } from "../hooks/useBoardData";
import { useCreateGame } from "../hooks/useCreateGame";
import {Notification} from "../componetns/Notification"
import { useState ,useEffect,useContext} from "react";
import * as apis from "../apis/api";
import { BoardData } from "../types/BoardData";
import './GamePage.css'
import { BoardDataContext } from "../providers/BoardDataProvider";

function GamePage(){

  const boardDataCtx=useContext(BoardDataContext)

  const onClick = () =>{
      apis.postBoardData().then((boardData:BoardData)=>{
        boardDataCtx.setBoardId(boardData.data.BoardId)
      })
  }
  
  // console.log("boardDataCtx.boardId:GamePage",boardDataCtx.boardId)
  // console.log("boardDataCtx.isMyTurn:GamePage",boardDataCtx.isMyTurn)
  let {stonesPos} = useBoardData(boardDataCtx.boardId,boardDataCtx.isMyTurn)

  return (
    <>
    {stonesPos.length ?
      <Notification isMyTurn={boardDataCtx.isMyTurn}>
        <Grid stonesPos={stonesPos}/>
      </Notification>
      :
      <div>
          <h1>オセロゲーム</h1>
          <div className="button" onClick={() => onClick()}>新規ゲームを始める</div>
      </div>
    }
    </>
  )

}


export default GamePage;
