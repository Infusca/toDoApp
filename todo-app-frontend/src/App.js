import logo from './logo.svg';
import './App.css';
// include todo component (app.js file in components)
import ToDo from './components/toDo';

function App() {
  return (
    <div className="App">
      <ToDo />
    </div>
  );
}

export default App;
