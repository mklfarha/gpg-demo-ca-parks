import React from 'react'
import { useParams } from 'react-router-dom';
import Header from '../components/Header';
import Client from '../client'
import {gql} from '@apollo/client';
import ParkForm from '../forms/ParkForm';
import { usePopupManager } from "react-popup-manager";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { v4 as uuidv4 } from 'uuid';
import { Link } from "react-router-dom";
import { MdCreate } from "react-icons/md";


class ParkEdit extends React.Component {
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
                query ParkByID($id: ID!) {
                    ParkByID(id:$id){ 
                        id,
                        name,
                        main_image,
                        phone,
                        hours,
                        allows_dogs,
                        links{
                            type,link,
                        },
                        status,
                        created_at,
                        updated_at,
                        recreation_area_id,
                        user_id,
                    }
                }
                `,
                variables:{
                    id:this.state.entity.id,                    
                }
            })
            .then((result) => {
                this.state.pending = false
                var data = result.data.ParkByID
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
                mutation upsertPark($input: ParkInput!) {
                    upsertPark(input:$input){ 
                        id
                    }
                }
                `,
                variables:{
                    input:{
                        id:(this.state.entity.id ? this.state.entity.id : ''),
                        name:this.state.entity.name ? this.state.entity.name : '',
                        main_image:this.state.entity.main_image ? this.state.entity.main_image : '',
                        phone:this.state.entity.phone ? this.state.entity.phone : '',
                        hours:this.state.entity.hours ? this.state.entity.hours : '',
                        allows_dogs:this.state.entity.allows_dogs ? this.state.entity.allows_dogs : false,
                        links:this.state.entity.links ? this.state.entity.links : [],
                        status:this.state.entity.status ? this.state.entity.status : '',
                        created_at:this.state.entity.created_at ? this.state.entity.created_at : '1999-09-09 09:09:09',
                        updated_at:this.state.entity.updated_at ? this.state.entity.updated_at : '1999-09-09 09:09:09',
                        recreation_area_id:this.state.entity.recreation_area_id ? this.state.entity.recreation_area_id : '',
                        user_id:this.state.entity.user_id ? this.state.entity.user_id : '',
                    }                    
                }
            })
            .then((result) => {
                var data = result.data.upsertPark
                if (!data.id) {
                    this.state.notFound = true
                } else {                    
                    this.state.entity.id = data.id
                    window.history.replaceState(null, document.title, "/park/edit/"+this.state.entity.id)
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
                    <h2 className="pageTitle">{this.props.params.id === undefined ? 'Create' : 'Edit'} Park</h2>                                  
                    { this.props.params.id !== undefined ? 
                    <a onClick={() => {
                        window.history.replaceState(null, document.title, "/park/edit/")
                        window.location.reload(false);
                    }} ><MdCreate /> Create New</a>
                    : undefined }
                </div>
                {this.state.pending ? <div className="formWrapper" style={ {textAlign:"center"} }>Loading...</div> : 
                    <ParkForm {...this.props} entity={this.state.entity} onSubmit={(data) =>{
                        this.upsertEntity()                        
                    } } />
                }
                <ToastContainer />
            </div> 
        )
    }
}

export default (props) => <ParkEdit {...props}
            params={useParams()} popups={usePopupManager()} />