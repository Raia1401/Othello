import './drawSquare.css'

function drawSquare(posX:number,posY:number,colorOfStone:number){
  const handleClick = () =>{
    console.log(posX,posY)
  }
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
        <div className="Square Click" onClick={handleClick}></div>
    }
    </>
  )
}

export default drawSquare;
