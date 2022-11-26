import './Grid.css'
import DrawSquare from './DrawSquare'



const Grid:React.FC<{stonesPos:number[][]}>=(props)=>{

  const BOARD_SIZE:number=8
  const grid=[]
  let row=[]

  for (let x=0; x< BOARD_SIZE; x++){
    for (let y=0; y< BOARD_SIZE; y++){
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
