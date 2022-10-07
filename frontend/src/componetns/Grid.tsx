import './Grid.css'
import drawSquare from './drawSquare'

function Grid(props:{stonesPosition:number[][]}){
  const BOARD_SIZE:number =5
  const grid=[]
  var rows=[]
  for (var x=0; x< BOARD_SIZE; x++){
    for (var y=0; y< BOARD_SIZE; y++){
      var stoneColor = props.stonesPosition[x][y]
      rows.push(drawSquare(x,y,stoneColor))
    }
    grid.push(<div className='Grid-row'>{rows}</div>)
    rows=[]
  }

  return (
    <div className='Grid'>{grid}</div>
  )
}

export default Grid;
