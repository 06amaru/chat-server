import './App.css';
import {ChakraProvider} from '@chakra-ui/react'
import Chat from './chat/Chat';


function App() {
  return (
    <div className="App">
      <ChakraProvider>
        <header className="App-header">
          <p>
            I ❤️ React
          </p>
        </header>
        <Chat />
      </ChakraProvider>
    </div>
  );
}

export default App;
