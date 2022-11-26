import './DrawSquare.css'
import { useStoneMovement } from '../hooks/useStoneMovement';

type Square ={
  x:number,
  y:number,
  colorOfStone:number
}

const myColor=1

const DrawSquare:React.FC<Square>=(props)=>{

  const {x,y,colorOfStone}=props
  const onClickBoard = useStoneMovement()

  return (
    <>
    {colorOfStone ===1 ?
      <div className="Square">
        <div className="Square-stone Square-stone-black"></div>
      </div>
      :
      colorOfStone === 2 ?
        <div className="Square">
          <div className="Square-stone Square-stone-white"></div>
        </div>
        :
        <div className="Square Click" onClick={() => onClickBoard(x+1,y+1,myColor)}></div>
    }
    </>
  )
}

export default DrawSquare;
