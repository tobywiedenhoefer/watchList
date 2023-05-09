import React from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import './style/App.css';
import WatchListComponent from "./WatchList";

function App(): JSX.Element {
  return (
      <Router>
        <div>
          <section>
            <Routes>
              <Route path="/" element={<WatchListComponent />} />
            </Routes>
          </section>
        </div>
      </Router>
  );
}

export default App;
