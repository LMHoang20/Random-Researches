import { useSelector, useDispatch } from 'react-redux';
import {
  selectText,
  textSlice,
} from '../../slices/text/textSlice';

export default function GoodStructure() {
    return (
        <div>
            <h1>Good Structure</h1>
            <Layer index={4} />
            {Layer({index: 4})}
        </div>
    )
}

function Layer({index}) {
    const color = [(Math.random() * 120).toFixed(0), (Math.random() * 120).toFixed(0), (Math.random() * 120).toFixed(0)]
    return (
        <div style={{
            backgroundColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
            padding: "1vh 1vw",
            margin: "1vh 1vw",
        }}>
            Layer {index}
            <br />
            {index > 1 ? <Layer index={index - 1} /> : 
                <TextLayer />
            }
        </div>
    )
}

function TextLayer() {
    const text = useSelector(selectText);
    const dispatch = useDispatch();
    const color = [(Math.random() * 120).toFixed(0), (Math.random() * 120).toFixed(0), (Math.random() * 120).toFixed(0)]
    return (
        <div style={{
            backgroundColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
            padding: "1vh 1vw",
            margin: "1vh 1vw"}}
            >
            <textarea value={text} onChange={(e) => dispatch(textSlice.actions.setText(e.target.value))}
                style={{
                    height: "50vh",
                    width: "50vw",
                }}
            />
        </div>
    )
}