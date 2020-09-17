import React, {useState} from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Nav from './Components/Nav/Nav.js'
import Home from './Components/Home/Home.js'
import Category from './Components/Category/Category.js'
import Admin from './Components/Admin/Admin.js'
import About from './Components/About.js'
import Footer from './Components/Footer.js'
import ProductsTable from './Components/Admin/ProductsTable/ProductsTable.js'
import CategoriesTable from './Components/Admin/CategoriesTable/CategoriesTable.js'
import StoresTable from './Components/Admin/StoresTable/StoresTable.js'


const NotFound = () => {
  return (<div>404</div>)
}
function App() {
  const [search, setSearch] = useState([])

  const handleSearch = (_search) => {
    setSearch(_search)
  }
  return (
    <Router>
      <Nav search={search} >
        <Switch>
          <Route exact path="/">
            <Home handleSearch={handleSearch} />
          </Route>
          <Route exact path="/about" component={About} />
          <Route exact path="/admin" component={Admin} />
          <Route exact path="/admin/products" component={ProductsTable} />
          <Route exact path="/admin/categories" component={CategoriesTable} />
          <Route exact path="/admin/stores" component={StoresTable} />
          <Route exact path="/:categ" >
            <Category handleSearch={handleSearch} />
          </Route>
          <Route path="/" component={NotFound} />
        </Switch>
      </Nav>
      <Footer title={`Made with love by Youssef Siam`} />
    </Router>
  );
}

export default App;
