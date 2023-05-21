import './grid.css'
import DrawSquare from './drawSquare'
import * as stttings from "../settings/settings"



const Grid:React.FC<{stonesPos:number[][]}>=(props)=>{

  const grid=[]
  let row=[]

  for (let x=0; x< stttings.BOARD_SIZE; x++){
    for (let y=0; y< stttings.BOARD_SIZE; y++){
      let stoneColor:number =props.stonesPos[x][y]
      let keyName:string = 'square:'+x.toString()+'_'+y.toString()
      row.push(<DrawSquare x={x} y={y} colorOfStone={stoneColor} key={keyName}/>)
    }
    let keyName:string = 'row:'+x.toString()
    grid.push(<div className='Grid-row' key={keyName}>{row}</div>)
    row=[]
  }

  return (
    <div className='Grid' key={111} >{grid}</div>
  )
}

export default Grid;
