import { Route, createBrowserRouter, createRoutesFromElements } from 'react-router-dom';
import Home from '../pages/home';
import BadStructure from '../pages/badStructure';
import GoodStructure from '../pages/goodStructure';

export const router = createBrowserRouter(
    createRoutesFromElements(
        <Route>
            <Route path="/" element={<Home/>} />
            <Route path="/badStructure" element={<BadStructure/>} />
            <Route path="/goodStructure" element={<GoodStructure/>} />
        </Route>
    )
)