import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import Login from './pages/Login';
import {
  BrowserRouter,
  Routes,
  Route,  
} from "react-router-dom";
import { PopupProvider } from "react-popup-manager";


import Park from "./routes/Park";
import ParkEdit from "./routes/ParkEdit";

import RecreationArea from "./routes/RecreationArea";
import RecreationAreaEdit from "./routes/RecreationAreaEdit";

import User from "./routes/User";
import UserEdit from "./routes/UserEdit";

import Feature from "./routes/Feature";
import FeatureEdit from "./routes/FeatureEdit";

import ParkHasFeature from "./routes/ParkHasFeature";
import ParkHasFeatureEdit from "./routes/ParkHasFeatureEdit";

import Event from "./routes/Event";
import EventEdit from "./routes/EventEdit";


const root = ReactDOM.createRoot(document.getElementById('root')); 
root.render(   
    <PopupProvider>    
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<App />} />
            <Route path="/login" element={<Login />} />
            
            <Route path="park" element={<Park />} />
            <Route path="park/edit" element={<ParkEdit />} />
            <Route path="park/edit/:id" element={<ParkEdit />} />
            
            <Route path="recreation_area" element={<RecreationArea />} />
            <Route path="recreation_area/edit" element={<RecreationAreaEdit />} />
            <Route path="recreation_area/edit/:id" element={<RecreationAreaEdit />} />
            
            <Route path="user" element={<User />} />
            <Route path="user/edit" element={<UserEdit />} />
            <Route path="user/edit/:id" element={<UserEdit />} />
            
            <Route path="feature" element={<Feature />} />
            <Route path="feature/edit" element={<FeatureEdit />} />
            <Route path="feature/edit/:id" element={<FeatureEdit />} />
            
            <Route path="park_has_feature" element={<ParkHasFeature />} />
            <Route path="park_has_feature/edit" element={<ParkHasFeatureEdit />} />
            <Route path="park_has_feature/edit/:id" element={<ParkHasFeatureEdit />} />
            
            <Route path="event" element={<Event />} />
            <Route path="event/edit" element={<EventEdit />} />
            <Route path="event/edit/:id" element={<EventEdit />} />
                    
        </Routes>
    </BrowserRouter>
    </PopupProvider>
);
