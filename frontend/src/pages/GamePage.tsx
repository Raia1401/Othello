import Grid from "../componetns/Grid";
import { useBoardData } from "../hooks/useBoardData";
import {Loding} from "../componetns/Loding"

function GamePage(){
  // let all = useBoardData()
  // let stonesPos=all[2]

  let {boardId,isMyTurn,stonesPos} = useBoardData()
  console.log("let isMyTurn:",isMyTurn)

  return (
    <>
    {stonesPos.length ?
      isMyTurn ?
        <Grid stonesPos={stonesPos}/>
        :
        <Loding>
          <Grid stonesPos={stonesPos}/>
        </Loding>
      :
      <h1>Wait a seconds...</h1>
    }
    </>
  )

}


export default GamePage;
