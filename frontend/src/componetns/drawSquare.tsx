import './drawSquare.css'
import { useStoneMovement } from '../hooks/useStoneMovement';
import  * as settings from "../settings/settings"

type Square ={
  x:number,
  y:number,
  colorOfStone:number
}

const DrawSquare:React.FC<Square>=(props)=>{

  const {x,y,colorOfStone}=props
  const onClickBoard = useStoneMovement()

  return (
    <>
    {colorOfStone ===settings.MY_STONE_COLOR ?
      <div className="Square">
        <div className="Square-stone Square-stone-black"></div>
      </div>
      :
      colorOfStone === settings.OPPONENT_STONE_COLOR ?
        <div className="Square">
          <div className="Square-stone Square-stone-white"></div>
        </div>
        :
        <div className="Square Click" onClick={() => onClickBoard(x+1,y+1,settings.MY_STONE_COLOR)}></div>
    }
    </>
  )
}

export default DrawSquare;
