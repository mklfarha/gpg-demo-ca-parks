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



    

    


class ParkLinksForm extends React.Component {
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
            <div className="formWrapper"  style={{borderTopColor:StringToColour('park_links')}}> 
                    <Formik    
                        initialValues={ this.state.entity }                    
                        validate={values => {
                            const errors = {};   
                            
                            if (!this.state.entity.type) {
                            
                                errors.type = "Required";
                            
                            }
                            if (!this.state.entity.link) {
                            
                                errors.link = "Required";
                            
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
                                toastId: "park_links-validation"
                            });                                
                        } else {
                            toast.dismiss("park_links-validation");
                        }
                        return (                            
                            <form onSubmit={handleSubmit}>
                                
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Type*</span>
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                <Input
                                    type="select"                                    
                                    name="type"
                                    onChange={(e) => {
                                        this.state.entity.type = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.type ? this.state.entity?.type : '' }
                                    
                                >
                                    <option value=""></option>
                                    
                                    <option value="brochure">Brochure</option>
                                    
                                    <option value="map">Map</option>
                                    
                                    <option value="instagram">Instagram</option>
                                    
                                    <option value="facebook">Facebook</option>
                                    
                                    <option value="youtube">Youtube</option>
                                    
                                    <option value="reservation">Reservation</option>
                                    
                                    <option value="chargers">Charger Locations</option>
                                    
                                    <option value="ebikes">E-Bike Locations</option>
                                    
                                </Input>
                                
                                

                                
                                
            
                                { errors.type ? <span className="info">{errors.type}</span> : undefined }
                                </FormGroup>
                                                                
                                <FormGroup className="inputWrapper">
                                <span>Link*</span>
                                
                                
                                
                                <Input
                                    type="text"
                                    name="link"
                                    onChange={(e) => {
                                        this.state.entity.link = e.target.value
                                        this.setState(this.state)
                                    } }                                        
                                    onBlur={handleBlur}
                                    value={this.state.entity?.link ? this.state.entity?.link : '' }
                                    
                                />
                                
                                
                                
                                                                    
                                
                                
                                
                                
                                

                                
                                
            
                                { errors.link ? <span className="info">{errors.link}</span> : undefined }
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

export default ParkLinksForm