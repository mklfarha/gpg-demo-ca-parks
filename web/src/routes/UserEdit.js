import React from 'react'
import { useParams } from 'react-router-dom';
import Header from '../components/Header';
import Client from '../client'
import {gql} from '@apollo/client';
import UserForm from '../forms/UserForm';
import { usePopupManager } from "react-popup-manager";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { v4 as uuidv4 } from 'uuid';
import { Link } from "react-router-dom";
import { MdCreate } from "react-icons/md";


class UserEdit extends React.Component {
    constructor(props) {
        super(props)        
        this.state = {
            pending:false,
            entity:{
                id: this.props.params.id,
            },
            notFound:false,
            images:[]
        }                             
    }

    componentDidMount() {        
        if (this.props.params.id !== undefined) {
            this.state.pending = true
            this.setState(this.state)
            this.fetchEntity();
        }
    }

    fetchEntity() {
        Client
            .query({
                query: gql`
                query UserByID($id: ID!) {
                    UserByID(id:$id){ 
                        id,
                        name,
                        email,
                        password,
                        status,
                        created_at,
                        updated_at,
                    }
                }
                `,
                variables:{
                    id:this.state.entity.id,                    
                }
            })
            .then((result) => {
                this.state.pending = false
                var data = result.data.UserByID
                if (data.length === 0) {
                    this.state.notFound = true
                    toast.error('Entity not found after insert',{
                        position: toast.POSITION.TOP_CENTER
                    });
                } else {
                    console.log(data[0])
                    this.state.entity = JSON.parse(JSON.stringify(data[0]))
                }                              
                this.setState(this.state)
            })
            .catch((err) => {
                toast.error('Error fetching entity',{
                        position: toast.POSITION.TOP_CENTER
                });
                console.error(err);
            });
    }

    upsertEntity() {
        Client
            .mutate({
                mutation: gql`
                mutation upsertUser($input: UserInput!) {
                    upsertUser(input:$input){ 
                        id
                    }
                }
                `,
                variables:{
                    input:{
                        id:(this.state.entity.id ? this.state.entity.id : ''),
                        name:this.state.entity.name ? this.state.entity.name : '',
                        email:this.state.entity.email ? this.state.entity.email : '',
                        password:this.state.entity.password ? this.state.entity.password : '',
                        status:this.state.entity.status ? this.state.entity.status : '',
                        created_at:this.state.entity.created_at ? this.state.entity.created_at : '1999-09-09 09:09:09',
                        updated_at:this.state.entity.updated_at ? this.state.entity.updated_at : '1999-09-09 09:09:09',
                    }                    
                }
            })
            .then((result) => {
                var data = result.data.upsertUser
                if (!data.id) {
                    this.state.notFound = true
                } else {                    
                    this.state.entity.id = data.id
                    window.history.replaceState(null, document.title, "/user/edit/"+this.state.entity.id)
                }                 
                window.scrollTo(0, 0)
                toast.success(this.props.params.id == undefined ? 'Created Successfully' : 'Edited Successfully',{
                    position: toast.POSITION.TOP_CENTER
                });
                setTimeout(() => {
                    toast.dismiss()
                    window.location.reload(false);
                }, "500")
                
            })
            .catch((err) => {
                toast.error('Error upserting entity',{
                        position: toast.POSITION.TOP_CENTER
                });
                console.error(err);
            });
    }    

    render() {              
        return (            
           <div>
                <Header />  
                <div className="actionsBar">  
                    <h2 className="pageTitle">{this.props.params.id === undefined ? 'Create' : 'Edit'} User</h2>                                  
                    { this.props.params.id !== undefined ? 
                    <a onClick={() => {
                        window.history.replaceState(null, document.title, "/user/edit/")
                        window.location.reload(false);
                    }} ><MdCreate /> Create New</a>
                    : undefined }
                </div>
                {this.state.pending ? <div className="formWrapper" style={ {textAlign:"center"} }>Loading...</div> : 
                    <UserForm {...this.props} entity={this.state.entity} onSubmit={(data) =>{
                        this.upsertEntity()                        
                    } } />
                }
                <ToastContainer />
            </div> 
        )
    }
}

export default (props) => <UserEdit {...props}
            params={useParams()} popups={usePopupManager()} />