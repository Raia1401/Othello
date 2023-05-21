import './grid.css'
import * as settings from "../settings/settings"
import './resultNotification.css'


const ResultNotification:React.FC<{stonesPos:number[][]}>=(props)=>{

  let myStoneNum:number=0
  let opponentStoneNum:number=0

  for (let x=0; x< settings.BOARD_SIZE; x++){
    for (let y=0; y< settings.BOARD_SIZE; y++){
      let stoneColor:number =props.stonesPos[x][y]
      if (stoneColor === settings.MY_STONE_COLOR){
        myStoneNum++
    }else if(stoneColor === settings.OPPONENT_STONE_COLOR){
        opponentStoneNum++
      }
    }
  }
  
  if(myStoneNum > opponentStoneNum){
    return (
      <div className='result-notification-win'>
        <h1>あなたの勝ちです🎉</h1>
        <p>あなたの石の数{myStoneNum}</p>
        <p>あなたの石の数{myStoneNum}</p>
        <p>相手の石の数{opponentStoneNum}</p>
      </div>
    )
  }else{
    return (
      <div className='result-notification-lose'>
        <h1>あなたの負けです...</h1>
        <p>あなたの石の数：{myStoneNum}</p>
        <p>相手の石の数：{opponentStoneNum}</p>
      </div>
    )
  }
}

export default ResultNotification;
