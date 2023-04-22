
import './App.css';
import { Link } from "react-router-dom";
import Header from "./components/Header";
import {StringToColour} from "./Utils";


function App() {    
  return (
    <div id="app" className="App">
      <Header />
      
      <Link to="/park">
        <div className="entityList" style={{borderTopColor:StringToColour('park')}}>
          Park
        </div>
      </Link>
      <Link to="/recreation_area">
        <div className="entityList" style={{borderTopColor:StringToColour('recreation_area')}}>
          RecreationArea
        </div>
      </Link>
      <Link to="/user">
        <div className="entityList" style={{borderTopColor:StringToColour('user')}}>
          User
        </div>
      </Link>
      <Link to="/feature">
        <div className="entityList" style={{borderTopColor:StringToColour('feature')}}>
          Feature
        </div>
      </Link>
      <Link to="/park_has_feature">
        <div className="entityList" style={{borderTopColor:StringToColour('park_has_feature')}}>
          ParkHasFeature
        </div>
      </Link>
      <Link to="/event">
        <div className="entityList" style={{borderTopColor:StringToColour('event')}}>
          Event
        </div>
      </Link>        
    </div>
  );
}

export default App;
