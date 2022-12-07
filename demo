=> Server.js

const express = require('express') 
const mongoos = require('mongoose') 
const cors = require('cors') 
const { default: mongoose } = require('mongoose') 
const app=express() 
//Routes:
const routesStud=require('./routes/student_route') 
const routeLogin = require('./routes/login') 
mongoose.connect('mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0/db1student') 
//Middlewares:
app.use(express.json()) 
app.use(cors()) 
app.use('/student',routesStud) 
app.use('/login',routeLogin) 
app.listen(5000,()=>{ 
 console.log('Server is running') 
}) 


================================================================================================================================

=>model.js


const mongoose = require('mongoose') 
const signup = new mongoose.Schema({ 
 name:{ 
 type:String, 
 required:true
 }, 
 email:{ 
 type:String, 
 required:true
 }, 
 phone:{ 
    type:String, 
    required:true
    }, 
    username:{ 
    type:String, 
    required:true
    }, 
    password:{ 
    type:String, 
    required:true
    }, 
    sign_date:{ 
    type:Date, 
    default:Date.now
    } 
   }) 
   module.exports=mongoose.model('student_master',signup) 


==========================================================================================================================

=>test.rest


POST http://localhost:5000/student/signup
content-type: application/json

{ 
 "name": "mukesh", 
 "email": "mukeshgmail.com", 
 "phone": "6352023139", 
 "username": "mukesh", 
 "password": "mukesh123"
} 
####
GET http://localhost:5000/student/dispStudent

###
POST http://localhost:5000/login/login
content-type: application/json
{
    "username": "mukesh", 
    "password": "mukesh123"
} 
###
DELETE http://localhost:5000/student/delete/6347297341d78c62de2d05cf

###
GET http://localhost:5000/student/getOneStudent/6346f4d5cb89428da59cab74

###
PATCH http://localhost:5000/student/update/6346f4d5cb89428da59cab74
content-type: application/json
{ 
 "name": "manish", 
 "email": "manish@gmail.com", 
 "phone": "9016941931", 
 "username": "manish", 
 "password": "manish123"
} 


============================================================================================================================

=>routes


const express= require('express') 
const router=express.Router() 
const signupinfo=require('../models/student_models') 
//INSERT
router.post('/signup', async(req,res)=>{ 
 const signupentry=new signupinfo({ 
 name:req.body.name, 
 email:req.body.email, 
 phone:req.body.phone, 
 username:req.body.username, 
 password:req.body.password, 
 }) 
 signupentry.save() 
}) 
//DISPLAY
router.get('/dispStudent',async(req,res)=>{ 
 const data= await signupinfo.find() 
 res.json(data) 
}) 
//Get data by ID
router.get('/getOneStudent/:id',async(req,res)=>{ 
    try{ 
    const data = await signupinfo.findById(req.params.id) 
    res.json(data) 
    }catch(e){ 
    res.json({message:e.message}) 
    } 
   }) 
   //Update 
   router.patch('/update/:id',async(req,res)=>{ 
    try{ 
    const id= req.params.id
    const updateData = req.body
    const options={new:true} 
    const result= await signupinfo.findByIdAndUpdate(id,updateData,options) 
    res.send(result) 
    }catch(e){ 
    res.json({message:e.message}) 
    } 
   }) 
   //Delete
   router.delete('/delete/:id',async(req,res)=>{ 
    try{ 
    const id=req.params.id
    const data= await signupinfo.findByIdAndDelete(id) 
    res.send('Student data has been deleted') 
    }catch(e){ 
    res.json({message:e.message}) 
    } 
   }) 
   //Search Record
   router.get('/search/:key',async(req,res)=>{ 
    let result= await signupinfo.find({ 
        "$or":[ 
        {name: {$regex: req.params.key}}, 
        {phone: {$regex: req.params.key}} 
        ] 
        }) 
        res.send(result) 
       }) 
       module.exports=router
       

=================================================================================================================================

=>login.js(server)


const express = require('express') 
const router = express.Router() 
const logininfo = require('../models/student_models') 
router.post('/login',async(req,res)=>{ 
 console.log(req.body) 
 if(req.body.username && req.body.password){ 
 let loginchk = await logininfo.findOne(req.body).select('-password') 
 if(loginchk){ 
 res.send(loginchk) 
 }
 else
 { 
 res.send({result:"Invalid username or password"}) 
 } 
 }
 else
 { 
 res.send({result:"Invalid username or password"}) 
 } 
}) 
module.exports=router


======================================================================================================================================

											CLIENT-SIDE

=======================================================================================================================================

=>App.js


import './App.css'; 
import { Route,Routes } from 'react-router-dom'; 
import CreateStudent from './comp/CreateStudent'; 
import DisplayStudent from './comp/DisplayStudents';
import Login from './comp/Login'; 
import Header from './comp/Header'; 
import Logout from './comp/Logout'; 
import UpdateStudent from './comp/UpdateStudent'; 
function App() { 
 return ( 
 <div className="App">
 <Header />
 <Routes>
 <Route path="/" element={<DisplayStudent/>} exact />
 <Route path="/AddStudents" element={<CreateStudent />} exact />
 <Route path="/UpdateStudent/:id" element={<UpdateStudent/>} exact />
 <Route path="/Login" element={ <Login />} exact />
 <Route path="/Logout" element={ <Logout />} exact />
 <Route path="*" element={ <DisplayStudent/>} exact/>
 </Routes>
 </div>
 ); 
} 
export default App; 


=====================================================================================================================================

=>CreateStudent.js

import React ,{useState}from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
const CreateStudent = () => { 
 const navigate = useNavigate() 
 const [name,setName]=useState('') 
 const [email,setEmail]=useState('') 
 const [phone,setPhone]=useState('') 
 const [username,setUsername]=useState('') 
 const [password,setPassword]=useState('') 
 const handleName=(e)=>{ 
 setName(e.target.value) 
 } 
 const handleEmail=(e)=>{ 
 setEmail(e.target.value) 
 } 
 const handlePhone=(e)=>{ 
 setPhone(e.target.value) 
 } 
 const handleUsername=(e)=>{ 
 setUsername(e.target.value) 
 } 
 const handlePassword=(e)=>{ 
 setPassword(e.target.value) 
 } 
 const collectData= async()=>{ 
 try{ 
 let insertData= axios.post('http://localhost:5000/student/signup',{ 
 name, 
 email, 
 phone, 
 username, 
 password
 }); 
 if(insertData){ 
 alert('Student Added Successfully.') 
 navigate('/',{replace:true}) 
 navigate(0) //To reload the page
 } 
 } 
 catch(error) { 
 console.log(error) 
 } 
 } 
 return ( 
 <div className='App'>
 <header className='App--header'>
 <h3>Student Registration</h3>
 <label>
 Name: 
 </label><br/>
 <input type="text" name="name" value={name} required
onChange={(e)=>handleName(e)}/><br/>
 <label>
 Email: 
 </label><br/>
 <input type="email" name="email" value={email} required
onChange={(e)=>handleEmail(e)}/><br/>
 <label>
 Phone: 
 </label><br/>
 <input type="text" name="phone" value={phone} required
onChange={(e)=>handlePhone(e)}/><br/>
 <label>
 Username: 
 </label><br/>
 <input type="text" name="username" value={username} required
onChange={(e)=>handleUsername(e)}/><br/>
 <label>
 Password: 
 </label><br/>
 <input type="password" name="password" value={password} required
onChange={(e)=>handlePassword(e)}/><br/>
 <button type="button" onClick={collectData}>Submit</button>
 </header>
 </div>
 ) 
} 
export default CreateStudent


===================================================================================================================================

=>DisplayStudents.js


import React,{useState,useEffect} from 'react'
import axios from 'axios'
import { NavLink } from 'react-router-dom'; 
const DisplayStudent = () => { 
 const [data,setData]=useState([]); 
 useEffect(()=>{ 
 getData() 
 },[]) 
 const getData= async()=>{ 
    try{ 
        const result = await axios.get('http://localhost:5000/student/dispStudent'); 
        setData(result.data) 
        }catch(e){ 
        console.log(e) 
        } 
        
        } 
        const deleteStudent=async(id)=>{ 
        alert(id) 
        let result = await axios.delete(`http://localhost:5000/student/delete/${id}`) 
        if(result){ 
        alert("Student data deleted") 
        getData() 
        } 
        } 
        const studentrecord = data.map((d)=>{ 
        return( 
        <tr>
        <td>{d.name}</td>
        <td>{d.email}</td>
        <td>{d.phone}</td>
        <td><NavLink to={`/UpdateStudent/${d._id}`}>Edit </NavLink></td> 
        <td><button type="button"
       onClick={()=>deleteStudent(d._id)}>Delete</button></td>
        </tr>
        ) 
        }) 
        const searchData=async(e)=>{ 
        let key = e.target.value
        if(key){ 
        let result = await axios.get(`http://localhost:5000/student/search/${key}`) 
        console.log(result.data) 
        if(result){ 
        setData(result.data) 
        } 
        }else{ 
        getData() 
        } 
        } 
        return ( 
        <div>
        Serach : <input type="text" onChange={searchData} />
        <table className='table table-stripped'>
        <tr>
 <th>Name</th>
 <th>Email</th>
 <th>Phone</th>
 <th>Edit</th>
 <th>Delete</th> 
 </tr>
 {studentrecord}
 </table>
 </div>
 ) 
} 
export default DisplayStudent


======================================================================================================================================

=>Login.js(client-side)


import React,{useState} from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
const Login = () => { 
 let navigate = useNavigate() 
 const [username,setUsername]=useState('') 
 const [password,setPassword]=useState('') 
 const handleUserChange=(e)=>{ 
 setUsername(e.target.value) 
 } 
 const handlePassChange=(e)=>{ 
 setPassword(e.target.value) 
 } 
 const collectData=async()=>{ 
 try{ 
 let logindata= await axios.post('http://localhost:5000/login/login',{ 
 username, 
 password
 }) 
 logindata=(await logindata).data
 alert(logindata.name) 
 if(logindata.name){ 
 localStorage.setItem("login",JSON.stringify(logindata)) 
 navigate('/',{replace:true}) 
 navigate(0) //To reload the page
 }else{
    alert("Username or password incorrect") 
 } 
 }catch(e){ 
 console.log(e) 
 } 
 } 
 return ( 
 <div className='App'>
 <header className='App--header'>
 <h1>Login Page</h1>
 <label>
 Username : 
 </label><br/>
 <input type="text" name="username" value={username} required
onChange={(e)=>handleUserChange(e)}/><br/>
 <label>
 Password : 
 </label><br/>
 <input type="password" name="password" value={password} required
onChange={(e)=>handlePassChange(e)}/><br/>
 <button type="button" onClick={collectData} >Login</button>
 </header>
 
 </div>
 ) 
} 
export default Login


======================================================================================================================================

=>Logout.js

import React,{useEffect} from 'react'
import { useNavigate } from 'react-router-dom'
const Logout = () => { 
 const navigate = useNavigate() 
 useEffect(()=>{ 
 localStorage.clear() 
 navigate('/',{replace:true}) 
 navigate(0) //To reload the page
 },[]) 
 return ( 
 <div>Logout</div>
 )
} 
export default Logout

=======================================================================================================================================

=>UpdateStudent.js


import React ,{useState,useEffect}from 'react'
import axios from 'axios'
import { useNavigate,useParams } from 'react-router-dom'
const UpdateStudent = () => { 
 const navigate = useNavigate() 
 const param = useParams() 
 const [name,setName]=useState('') 
 const [email,setEmail]=useState('') 
 const [phone,setPhone]=useState('') 
 const [username,setUsername]=useState('') 
 const [password,setPassword]=useState('') 
 
 const handleName=(e)=>{ 
 setName(e.target.value) 
 } 
 const handleEmail=(e)=>{ 
 setEmail(e.target.value) 
 } 
 const handlePhone=(e)=>{ 
 setPhone(e.target.value) 
 } 
 const handleUsername=(e)=>{ 
 setUsername(e.target.value) 
 } 
 const handlePassword=(e)=>{ 
 setPassword(e.target.value) 
 } 
 const dispData = async()=>{ 
 let result = await
axios.get(`http://localhost:5000/student/getOneStudent/${param.id}`) 
 setName(result.data.name) 
 setEmail(result.data.email) 
 setPhone(result.data.phone) 
 setUsername(result.data.username) 
 setPassword(result.data.password)
} 
useEffect(()=>{ //To be invoke on page load, invoked once use []
dispData() 
},[]) 
const collectData= async()=>{ 
let result = await axios.patch(`http://localhost:5000/student/update/${param.id}`,{ 
name, 
email, 
phone, 
username, 
password
}) 
if(result){ 
navigate('/') 
navigate(0) 
} 

} 
return ( 
<div className='App'>
<header className='App--header'>
<h3>Update Student</h3>
<label>
Name: 
</label><br/>
<input type="text" name="name" value={name} required
onChange={(e)=>handleName(e)}/><br/>
<label>
Email: 
</label><br/>
<input type="email" name="email" value={email} required
onChange={(e)=>handleEmail(e)}/><br/>
<label>
Phone: 
</label><br/>
<input type="text" name="phone" value={phone} required
onChange={(e)=>handlePhone(e)}/><br/>
<label>
Username: 
</label><br/>
<input type="text" name="username" value={username} required
onChange={(e)=>handleUsername(e)}/><br/>
<label>
Password: 
</label><br/>
<input type="password" name="password" value={password} required
onChange={(e)=>handlePassword(e)}/><br/>
 <button type="button" onClick={collectData}>Update</button>
 </header>
 
 </div>
 ) 
} 
export default UpdateStudent

======================================================================================================================================

=>Header.js:

import React from 'react'
import { NavLink } from 'react-router-dom'
const Header = () => { 
 let auth= localStorage.getItem('login') 
 return ( 
 <div>
 <nav>
 <NavLink to='/'>Home</NavLink>
 {
 auth?<>
 <NavLink to='/Logout'>Logout</NavLink>
 <NavLink to='/AddStudents'>Add Student</NavLink> 
 </>: 
 <NavLink to='/Login'>Login</NavLink>
 }
 </nav>
 </div>
 ) 
} 
export default Header

=======================================================================================================================================

=>App.css: (just Append) 

nav{ 
  background-color: #333; 
  padding: 5px; 
} 
nav a{ 
 margin: 20px; 
} 
nav a.active{ 
 font-weight: bold; 
 text-decoration: none; 
 color:beige; 
}    


======================================================================================================================================

=> index.js (just Append) 


import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter } from 'react-router-dom'; 

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  
  <BrowserRouter>
  <App />
 </BrowserRouter>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();


======================================================================================================================================

=>Index.html (paste Bootstrap Links)

<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/css/bootstrap.min.css"
rel="stylesheet" integrity="sha384-
Zenh87qX5JnK2Jl0vWa8Ck2rdkQ2Bzep5IDxbcnCeuOxjzrPF/et3URy9Bv1WTRi" crossorigin="anonymous">

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/js/bootstrap.bundle.min.js"
integrity="sha384-OERcA2EqjJCMA+/3y+gxIOqMEjwtxJY7qPCqsdltbNJuaOe923+mo//f6V8Qbsw3"
crossorigin="anonymous"></script

======================================================================================================================================

