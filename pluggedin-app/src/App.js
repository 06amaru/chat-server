import logo from './logo.svg';
import './App.css';
import { encrypt, decrypt, PrivateKey } from 'eciesjs'
import Chat from './chat/Chat'

function App() {

  const k1 = new PrivateKey()
  console.log(k1)
  const data = Buffer.from('jaoks')
  const encrypted = encrypt(k1.publicKey.toHex(), data)
  console.log(encrypted)
  const decrypted = decrypt(k1.toHex(), encrypted).toString()
  console.log(decrypted)
  return (
    <div>
      <Chat />
     </div>
  );
}

export default App;
