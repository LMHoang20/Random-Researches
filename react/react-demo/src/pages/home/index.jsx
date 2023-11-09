import { useInput } from '../../hooks/useInput'

function Home() {
    const [username, userInput] = useInput({ type: "text" });
    const [password, passwordInput] = useInput({ type: "text" });

    return (
        <div>
            <h1>Home</h1>
            <>
                {userInput} <br />
                {passwordInput} 
            </>
        </div>
    )
}

export default Home