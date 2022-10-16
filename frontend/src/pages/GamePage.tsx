import Grid from "../componetns/Grid";
import { useStonesPos } from "../hooks/useStonesPos";

function GamePage(){
  let stonesPos = useStonesPos()

  return (
    // <Grid stonesPosition={[[0,0,1,2,1],[0,0,1,2,1],[0,0,1,2,1],[0,0,1,2,1],[0,0,1,2,1]]}/>
    <Grid stonesPos={stonesPos}/>
  )

}


export default GamePage;
