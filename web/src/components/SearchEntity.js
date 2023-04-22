import React from 'react'
import AsyncSelect from 'react-select/async';
import Client from '../client'
import {gql} from '@apollo/client';
import { toast } from 'react-toastify';

export const searchQueries = new Map();
export const queriesByID = new Map();
export const searchFields = new Map();

searchQueries.set("park", gql`
                query SearchPark($query: String!, $limit: Int, $offset: Int) {
                    SearchPark(query:$query, limit:$limit, offset:$offset){ 
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
            `)
queriesByID.set("park", gql`
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
    `
)
searchFields.set("park", (item) => { var res = ''; res = item.name;return res;})
searchQueries.set("recreation_area", gql`
                query SearchRecreationArea($query: String!, $limit: Int, $offset: Int) {
                    SearchRecreationArea(query:$query, limit:$limit, offset:$offset){ 
                        id,
                        name,
                        status,
                        created_at,
                        updated_at,
                        user_id,
                    }
                }
            `)
queriesByID.set("recreation_area", gql`
    query RecreationAreaByID($id: ID!) {
                    RecreationAreaByID(id:$id){ 
                        id,
                        name,
                        status,
                        created_at,
                        updated_at,
                        user_id,
                    }
                }
    `
)
searchFields.set("recreation_area", (item) => { var res = ''; res = item.name;return res;})
searchQueries.set("feature", gql`
                query SearchFeature($query: String!, $limit: Int, $offset: Int) {
                    SearchFeature(query:$query, limit:$limit, offset:$offset){ 
                        id,
                        name,
                        status,
                        created_at,
                        updated_at,
                        user_id,
                    }
                }
            `)
queriesByID.set("feature", gql`
    query FeatureByID($id: ID!) {
                    FeatureByID(id:$id){ 
                        id,
                        name,
                        status,
                        created_at,
                        updated_at,
                        user_id,
                    }
                }
    `
)
searchFields.set("feature", (item) => { var res = ''; res = item.name;return res;})

export class SearchEntity extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            entityIdentifier: props.entityIdentifier,
            entityName: props.entityName,
            searchInput:"",
            value: props.value,            
        }   

        if (props.value?.value) {
            this.loadByID()
        }

        this.state.ref = React.createRef();
                 
    }

    loadByID(){
        if (!queriesByID.has(this.state.entityIdentifier)){
            console.log("query id not found : "+this.state.entityIdentifier)
            return 
        }

        Client
            .query({
                query: queriesByID.get(this.state.entityIdentifier),
                variables:{
                    id:this.state.value.value,                                       
                }
            })
            .then((result) => {  
                var data = result.data[this.state.entityName + "ByID"]
                if (data.length > 0) {
                    var fieldsFunc = searchFields.get(this.state.entityIdentifier)
                    this.state.value.label = fieldsFunc(data[0])
                    this.setState(this.state)
                }
            })

    }

    loadOptions(inputValue, callback){
        if (!searchQueries.has(this.state.entityIdentifier)){
            console.log("query not found : "+this.state.entityIdentifier)
            return 
        }
        Client
            .query({
                query: searchQueries.get(this.state.entityIdentifier),
                variables:{
                    query:this.state.searchInput,    
                    limit:10,
                    offset:0,                    
                }
            })
            .then((result) => {                
                var data = result.data["Search"+this.state.entityName]                
                var fieldsFunc = searchFields.get(this.state.entityIdentifier)
                var items = data.map(function(item) {
                    return {
                        value:item.id,
                        label:fieldsFunc(item)
                    };
                });
                callback(items)                                                                 
            })
            .catch((err) => {
                toast.error('Error searching entity',{
                        position: toast.POSITION.TOP_CENTER
                });
                console.error(err);
            });
    }

    handleInputChange(newValue){        
        this.state.searchInput = newValue
        this.setState(this.state)
        return newValue;
    };

    clearValue() {
        this.state.value = ""
        this.setState(this.state)
    }

    render() {
        return (
            <div>                
                <AsyncSelect
                    onChange={(data) => {
                        this.state.value = data
                        this.props.onChange(data, this.state.entityIdentifier)
                    } }
                    placeholder={this.props.placeholder}
                    isClearable={true}
                    escapeRemoves={true}                   
                    cacheOptions
                    loadOptions={(inputValue, callback) => this.loadOptions(inputValue, callback)}
                    defaultOptions
                    onInputChange={(newValue) => this.handleInputChange(newValue)}
                    value={this.state.value}                    
                    components={ {
                            IndicatorSeparator: () => null
                        } }
                    />
            </div>
        )
    }
}