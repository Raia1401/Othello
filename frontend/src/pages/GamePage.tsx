import Grid from "../componetns/Grid";
import { useStonesPos } from "../hooks/useStonesPos";

function GamePage(){
  let stonesPos = useStonesPos()
  // console.log("let stonesPos:",stonesPos)

  return (
    <>
    { stonesPos.length ?
      <Grid stonesPos={stonesPos}/>
      :
      <h1>Now Loading...</h1>
    }
    </>
  )

}


export default GamePage;
