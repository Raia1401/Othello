import {useContext} from "react";

import './gamePage.css'
import Grid from "../componetns/grid";
import { useBoardData } from "../hooks/useBoardData";
import { BoardDataContext } from "../providers/boardDataProvider";
import ResultNotification from "../componetns/resultNotification";
import TurnNotification from "../componetns/turnNotification"
import { useNewGame } from "../hooks/useNewGame";
import { useTurnPass } from "../hooks/useTurnPass";

function GamePage(){

  const boardDataCtx=useContext(BoardDataContext)
  let {stonesPos,isMatchEnd} = useBoardData(boardDataCtx.boardId,boardDataCtx.isMyTurn)
  const createGame=useNewGame()
  const passMyTurn=useTurnPass()
  
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
          <div　className="pass_button" onClick={() => passMyTurn()}>パスをする</div>
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
