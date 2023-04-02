import Grid from "../componetns/Grid";
import { useBoardData } from "../hooks/useBoardData";
import {Notification} from "../componetns/Notification"
import {useContext} from "react";
import * as apis from "../apis/api";
import { BoardData } from "../types/BoardData";
import './GamePage.css'
import { BoardDataContext } from "../providers/BoardDataProvider";

function GamePage(){

  const boardDataCtx=useContext(BoardDataContext)

  const createGame = () =>{
      apis.postBoardData().then((boardData:BoardData)=>{
        boardDataCtx.setBoardId(boardData.data.BoardId)
      })
  }
  let {stonesPos,isMatchEnd} = useBoardData(boardDataCtx.boardId,boardDataCtx.isMyTurn)
  // let {stonesPos,isMatchEnd} = useBoardData(81,boardDataCtx.isMyTurn)

  return (
    <div className="contents">
    {isMatchEnd ? 
      <div>終了です</div>
      :
      stonesPos.length ?
        <>
          <Notification isMyTurn={boardDataCtx.isMyTurn}>
            <Grid stonesPos={stonesPos}/>
          </Notification>
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
