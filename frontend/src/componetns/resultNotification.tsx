import './grid.css'
import * as settings from "../settings/settings"


const ResultNotification:React.FC<{stonesPos:number[][]}>=(props)=>{

  const BOARD_SIZE:number=8
  let myStoneNum:number=0
  let opponentStoneNum:number=0

  for (let x=0; x< BOARD_SIZE; x++){
    for (let y=0; y< BOARD_SIZE; y++){
      let stoneColor:number =props.stonesPos[x][y]
      if (stoneColor === settings.myColor){
        myStoneNum++
    }else if(stoneColor === settings.opponentColor){
        opponentStoneNum++
      }
    }
  }
  
  if(myStoneNum > opponentStoneNum){
    return (
      <div className='result-notification'>
        <p>あなたの勝ちです</p>
      </div>
    )
  }else{
    return (
      <div className='result-notification'>
        <p>あなたの負けです</p>
      </div>
    )
  }
}

export default ResultNotification;
