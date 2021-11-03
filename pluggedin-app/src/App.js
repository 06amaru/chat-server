import logo from './logo.svg';
import './App.css';
import { encrypt, decrypt, PrivateKey } from 'eciesjs'

function App() {

  const k1 = new PrivateKey()
  console.log(k1)
  const data = Buffer.from('jaoks')
  const encrypted = encrypt(k1.publicKey.toHex(), data)
  console.log(encrypted)
  const decrypted = decrypt(k1.toHex(), encrypted).toString()
  console.log(decrypted)
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
