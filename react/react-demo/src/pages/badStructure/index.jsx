import React, { useState } from 'react';

export default function BadStructure() {
    const [text, setText] = useState("");
    return (
        <div>
            <h1>Bad Structure</h1>
            <Layer index={1} text={text} setText={setText} />
            {Layer({index: 1, text: text, setText: setText})}
        </div>
    )
}

function Layer({index, text, setText}) {
    const color = [(Math.random() * 120).toFixed(0), (Math.random() * 120).toFixed(0), (Math.random() * 120).toFixed(0)]
    return (
        <div style={{
            backgroundColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
            padding: "1vh 1vw",
            margin: "1vh 1vw",
        }}>
            Layer {index}
            <br />
            {index < 4 ? <Layer index={index + 1} text={text} setText={setText} /> : 
                <textarea value={text} onChange={(e) => setText(e.target.value)}
                    style={{
                        backgroundColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
                        height: "50vh",
                        width: "50vw",
                    }}
                />
            }
        </div>
    )
}