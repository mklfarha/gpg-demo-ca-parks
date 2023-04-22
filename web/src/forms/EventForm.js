import React from 'react'
import { Formik } from 'formik';
import { Card, CardBody, CardTitle, CardText, Form, Label, Input, Button, FormGroup } from 'reactstrap'
import ImageUploading from 'react-images-uploading';
import { GPGModal } from '../components/GPGModal';
import { toast } from 'react-toastify';
import axios from 'axios';
import {StringToColour} from "../Utils";
import FetchUser from '../FetchUser';
import {SearchEntity} from '../components/SearchEntity'



    

    

    

    

    

    

    

    

    

    

    


class EventForm extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            entity: props.entity
        }           

        
                            
                                
            
        
                            
                                
            
        
            
        
                            
                                
            
        
                            
                                
            
        
                            
                                
            
        
                            
                                
            
        
                            
                                
            
        
                            
                                
            
        
                            
                                
            
        
                            
                                
            
        
    }

    componentDidMount() {
        FetchUser().then(res => {
            if (res.status === 200) {
                this.state.user = res.data
                this.setState(this.state)
            }
        })
    }

    onImageChange(imageList, addUpdateIndex, fieldIdentifier) {
        // data for submit
        console.log(imageList)
        if (imageList.length > 0) {
            var imageFile = imageList[0]
            var formData = new FormData();            
            formData.append("file", imageFile.file);
            axios.post('http://localhost:8080/upload', formData, {
                headers: {
                'Content-Type': 'multipart/form-data'
                }
            }).then(response => {                    
                    if (response && response.data && response.data.file_path) {
                        this.state.entity[fieldIdentifier] = response.data.file_path
                        this.setState(this.state)
                    }
                }
            )
        }
    };

    render() {
        return (
            <div className="formWrapper"  style={{borderTopColor:StringToColour('event')}}> 
                    <Formik    
                        initialValues={ this.state.entity }                    
                        validate={values => {
                            const errors = {};   
                            
                            if (!this.state.entity.name) {
                            
                                errors.name = "Required";
                            
                            }
                            if (!this.state.entity.main_image) {
                            
                                errors.main_image = "Required";
                            
                            }
                            if (!this.state.entity.start_date) {
                            
                                errors.start_date = "Required";
                            
                            }
                            if (!this.state.entity.end_date) {
                            
                                errors.end_date = "Required";
                            
                            }
                            if (!this.state.entity.status) {
                            
                                errors.status = "Required";
                            
                            }
                            if (!this.state.entity.created_at) {
                            
                            }
                            if (!this.state.entity.updated_at) {
                            
                            }
                            if (!this.state.entity.park_id) {
                            
                                errors.park_id = "Required";
                            
                            }
                            if (!this.state.entity.user_id) {
                            
                                errors.user_id = "Required";
                            
                            }                                               
                            return errors;
                        }}  
                        onSubmit={(values, actions) =>{                            
                            this.props.onSubmit(this.state.entity)                                
                        }}                      
                        >
                        {({                            
                            errors,   
                            handleBlur,
                            handleSubmit,
                            isValid,                                                                                 
                        }) => {
                        if (!isValid) {                                
                            toast.error('Error validating form',{
                                position: toast.POSITION.TOP_CENTER,
                                toastId: "event-validation"
                            });                                
                        } else {
                            toast.dismiss("event-validation");
                        }
                        return (                            
                            <form onSubmit={handleSubmit}>
                                
                                                                
                                <FormGroup className="inputWrapper">
                                <span>ID*</span>
                                
                                
                                <Input
                                    type="text"
                                    name="id"
                                    onChange={(e) => {
                                        this.state.entity.id = e.target.value
                                        this.setState(this.state)
                                    } }
                                    onBlur={handleBlur}                                        
                                    value={this.state.entity?.id ? this.state.entity?.id : '' }
                                    disabled                                  
                                />
                                
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.id ? <span className="info">{errors.id}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Name*</span>
                                
                                
                                
                                <Input
                                    type="text"
                                    name="name"
                                    onChange={(e) => {
                                        this.state.entity.name = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.name ? this.state.entity?.name : '' }
                                    
                                />
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.name ? <span className="info">{errors.name}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Description</span>
                                
                                
                                
                                <Input
                                    type="textarea"
                                    name="description"
                                    onChange={(e) => {
                                        this.state.entity.description = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.description ? this.state.entity?.description : '' }
                                    
                                />
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.description ? <span className="info">{errors.description}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Main Image*</span>
                                
                                
                                
                                <div>
                                <Input
                                    disabled
                                    type="text"
                                    name="main_image"
                                    onChange={(e) => {
                                        this.state.entity.main_image = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.main_image ? this.state.entity?.main_image : '' }                                    
                                />
                                <ImageUploading                                    
                                    value={ this.state.entity.main_image }
                                    onChange={(imageList, addUpdateIndex) => this.onImageChange(imageList, addUpdateIndex, "main_image")}
                                    onError={ (err) => {console.log(err)} }
                                    maxNumber={1}
                                    dataURLKey="data_url">
                                    {({
                                    imageList,
                                    onImageUpload,
                                    onImageRemoveAll,
                                    onImageUpdate,
                                    onImageRemove,
                                    isDragging,
                                    dragProps,
                                    }) => (
                                    // write your building UI
                                    <div className="upload__image-wrapper">
                                        <Button
                                        style={isDragging ? { color: 'red' } : undefined}
                                        onClick={onImageUpload}
                                        {...dragProps}
                                        >
                                        Click or Drop here
                                        </Button>
                                        &nbsp; 
                                        { this.state.entity.main_image ?                                        
                                        <div className="image-item">
                                            <img src={this.state.entity.main_image} alt="" width="100" />
                                            <div className="image-item__btn-wrapper">                                                
                                            <Button className="btn-danger" onClick={() => {
                                                this.state.entity.main_image = undefined
                                                this.setState(this.state)
                                            }}>Remove</Button>
                                            </div>
                                        </div>
                                        : '' }
                                    </div>
                                    )}
                                </ImageUploading>
                                </div>
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.main_image ? <span className="info">{errors.main_image}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Start Date*</span>
                                
                                
                                
                                                                    
                                
                                                                    
                                <Input
                                    type="date"
                                    name="start_date"
                                    onChange={(e) => {
                                        this.state.entity.start_date = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.start_date ? this.state.entity?.start_date : '' }                                  
                                    
                                />                                    
                                
                                
                                
                                

                                
                                
            
                                { errors.start_date ? <span className="info">{errors.start_date}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>End Date*</span>
                                
                                
                                
                                                                    
                                
                                                                    
                                <Input
                                    type="date"
                                    name="end_date"
                                    onChange={(e) => {
                                        this.state.entity.end_date = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.end_date ? this.state.entity?.end_date : '' }                                  
                                    
                                />                                    
                                
                                
                                
                                

                                
                                
            
                                { errors.end_date ? <span className="info">{errors.end_date}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Status*</span>
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                <Input
                                    type="select"                                    
                                    name="status"
                                    onChange={(e) => {
                                        this.state.entity.status = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.status ? this.state.entity?.status : '' }
                                    
                                >
                                    <option value=""></option>
                                    
                                    <option value="enabled">Enabled</option>
                                    
                                    <option value="disabled">Disabled</option>
                                    
                                </Input>
                                
                                

                                
                                
            
                                { errors.status ? <span className="info">{errors.status}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Created*</span>
                                
                                
                                
                                                                    
                                
                                
                                                                    
                                <Input
                                    type="text"
                                    name="created_at"
                                    onChange={(e) => {
                                        this.state.entity.created_at = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.created_at ? this.state.entity?.created_at : '' }
                                    disabled
                                />                                    
                                
                                
                                

                                
                                
            
                                { errors.created_at ? <span className="info">{errors.created_at}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Updated*</span>
                                
                                
                                
                                                                    
                                
                                
                                                                    
                                <Input
                                    type="text"
                                    name="updated_at"
                                    onChange={(e) => {
                                        this.state.entity.updated_at = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.updated_at ? this.state.entity?.updated_at : '' }
                                    disabled
                                />                                    
                                
                                
                                

                                
                                
            
                                { errors.updated_at ? <span className="info">{errors.updated_at}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Park*</span>
                                
                                
                                    
                                        <SearchEntity 
                                            entityIdentifier="park" 
                                            entityName="Park"
                                            onChange={(e) => {
                                                this.state.entity.park_id = e.value
                                                this.setState(this.state)                                                
                                            } }
                                            value={ {value:this.state.entity?.park_id, label:this.state.entity?.park_id ? "Loading..." : "Search"} }
                                             />
                                    
                                
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.park_id ? <span className="info">{errors.park_id}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>User ID*</span>
                                
                                                                
                                ({this.state.user?.name}) &nbsp;
                                {this.state.entity.user_id = this.state.user?.id }
                                
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.user_id ? <span className="info">{errors.user_id}</span> : undefined }
                                </FormGroup>
                                                                
                                
                                
                                <Button type="submit">
                                    Submit
                                </Button>
                            </form>
                        )}}
                        </Formik>                    
                </div>
        )
    }
}

export default EventForm