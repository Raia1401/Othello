import {useContext} from "react";

import './gamePage.css'
import Grid from "../componetns/grid";
import { useBoardData } from "../hooks/useBoardData";
import * as apis from "../apis/api";
import { BoardData } from "../types/boardData";
import { BoardDataContext } from "../providers/boardDataProvider";
import ResultNotification from "../componetns/resultNotification";
import TurnNotification from "../componetns/turnNotification"

function GamePage(){

  const boardDataCtx=useContext(BoardDataContext)

  const createGame = () =>{
      apis.postBoardData().then((boardData:BoardData)=>{
        boardDataCtx.setBoardId(boardData.data.BoardId)
      })
  }
  
  let {stonesPos,isMatchEnd} = useBoardData(boardDataCtx.boardId,boardDataCtx.isMyTurn)

  return (
    <div className="contents">
    {isMatchEnd ? 
      <ResultNotification stonesPos={stonesPos}/>
      :
      stonesPos.length ?
        <>
          <TurnNotification isMyTurn={boardDataCtx.isMyTurn}>
            <Grid stonesPos={stonesPos}/>
          </TurnNotification>
          <div　className="pass_button">パスをする</div>
        </>
        :
        <div>
            <h1>オセロゲーム</h1>
            <div className="newgame_button" onClick={() => createGame()}>新規ゲームを始める</div>
        </div>
    }
      
    </div>
  )

}


export default GamePage;
