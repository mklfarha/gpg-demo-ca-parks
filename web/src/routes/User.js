import React,{useMemo} from 'react'
import { useParams } from 'react-router-dom';
import Header from '../components/Header'
import DataTable from 'react-data-table-component';
import { MdCreate,MdOutlineViewWeek } from "react-icons/md";
import { Link } from "react-router-dom";
import Client,{CachedClient} from '../client'
import {gql} from '@apollo/client';
import {StringToColour} from "../Utils";
import { Input } from 'reactstrap'
import {SearchEntity, queriesByID, searchFields} from '../components/SearchEntity'
import { GPGModal } from '../components/GPGModal';
import { usePopupManager } from "react-popup-manager";

const columns = [
    
    {
        name: 'ID',        
        
        selector: row => row.id,
        
        
        width: "350px",
        
        wrap: true,        
        sortable: true,
        sortField: "id",  
        omit:false,      
    },
    {
        name: 'Name',        
        
        selector: row => row.name,
        
        
        width: "200px",
        
        wrap: true,        
        sortable: true,
        sortField: "name",  
        omit:false,      
    },
    {
        name: 'Email',        
        
        selector: row => row.email,
        
        
        width: "200px",
        
        wrap: true,        
        sortable: true,
        sortField: "email",  
        omit:false,      
    },
    {
        name: 'Password',        
        
        selector: row => row.password,
        
        
        width: "200px",
        
        wrap: true,        
        sortable: true,
        sortField: "password",  
        omit:false,      
    },
    {
        name: 'Status',        
        
        selector: row => row.status,
        
        
        width: "200px",
        
        wrap: true,        
        sortable: true,
        sortField: "status",  
        omit:false,      
    },
    {
        name: 'Created',        
        
        selector: row => row.created_at,
        
        
        width: "200px",
        
        wrap: true,        
        sortable: true,
        sortField: "created_at",  
        omit:false,      
    },
    {
        name: 'Created',        
        
        selector: row => row.updated_at,
        
        
        width: "200px",
        
        wrap: true,        
        sortable: true,
        sortField: "updated_at",  
        omit:false,      
    },    
];

var pressedKeys = {};

class User extends React.Component {
    constructor(props) {
        super(props)        
        this.state = { 
            pending:true,           
            items: [],
            pageSize:14,
            page:0,
            total: 10000,
            searchQuery:"",
            selectRefs: {},
            columns: [],            
        }

        
    }

    componentDidMount() {        
        this.fetchItems();

        this.state.columns = [
            
            {
                identifier: "id",
                name: 'ID',        
                
                selector: row => row.id,
                
                
                width: "350px",
                
                wrap: true,        
                sortable: true,
                sortField: "id",  
                omit:false,      
            },
            {
                identifier: "name",
                name: 'Name',        
                
                selector: row => row.name,
                
                
                width: "200px",
                
                wrap: true,        
                sortable: true,
                sortField: "name",  
                omit:false,      
            },
            {
                identifier: "email",
                name: 'Email',        
                
                selector: row => row.email,
                
                
                width: "200px",
                
                wrap: true,        
                sortable: true,
                sortField: "email",  
                omit:false,      
            },
            {
                identifier: "password",
                name: 'Password',        
                
                selector: row => row.password,
                
                
                width: "200px",
                
                wrap: true,        
                sortable: true,
                sortField: "password",  
                omit:false,      
            },
            {
                identifier: "status",
                name: 'Status',        
                
                selector: row => row.status,
                
                
                width: "200px",
                
                wrap: true,        
                sortable: true,
                sortField: "status",  
                omit:false,      
            },
            {
                identifier: "created_at",
                name: 'Created',        
                
                selector: row => row.created_at,
                
                
                width: "200px",
                
                wrap: true,        
                sortable: true,
                sortField: "created_at",  
                omit:false,      
            },
            {
                identifier: "updated_at",
                name: 'Created',        
                
                selector: row => row.updated_at,
                
                
                width: "200px",
                
                wrap: true,        
                sortable: true,
                sortField: "updated_at",  
                omit:false,      
            },    
        ]   

        var cachedColumns = localStorage.getItem("user_columns")     
        var parsedCachedColumns = JSON.parse(cachedColumns)        
        if (parsedCachedColumns !== undefined && parsedCachedColumns !== null) {            
            this.state.columns.map((column, index) => {                
                this.state.columns[index].omit = parsedCachedColumns[column.identifier]
            })
        }        
        this.setState(this.state)

        window.onkeyup = function(e) { pressedKeys[e.keyCode] = false; }
        window.onkeydown = function(e) { pressedKeys[e.keyCode] = true; }

    }

    fetchItems() {        
        const offset = this.state.page*this.state.pageSize        
        Client
            .query({
                query: gql`
                query UserByStatus($limit: Int, $offset: Int) {
                    UserByStatus(status:"enabled",limit:$limit, offset:$offset){ 
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
                    limit:this.state.pageSize,
                    offset:offset
                }
            })
            .then((result) => {
                
                const items = result.data.UserByStatus
                if (items.length < this.state.pageSize){
                    this.state.total = ((this.state.page)*this.state.pageSize)+items.length
                }                
                                       
                Promise.allSettled(EntityRefPromises(items)).then(()=>{
                    setTimeout(() => {                    
                        this.state.pending = false
                        this.state.items = items                        
                        this.setState(this.state)
                    }, 500)
                })
            });
            
    }

    searchItems() {        
        const offset = this.state.page*this.state.pageSize        
        Client
            .query({
                query: gql`
                query SearchUser($query:String!, $limit: Int, $offset: Int) {
                    SearchUser(query:$query,limit:$limit, offset:$offset){ 
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
                    query:this.state.searchQuery,
                    limit:this.state.pageSize,
                    offset:offset
                }
            })
            .then((result) => {
                
                const items = result.data.SearchUser
                if (items.length < this.state.pageSize){
                    this.state.total = ((this.state.page)*this.state.pageSize)+items.length
                }                
                Promise.allSettled(EntityRefPromises(items)).then(()=>{
                    setTimeout(() => {                    
                        this.state.pending = false
                        this.state.items = items                        
                        this.setState(this.state)
                    }, 500)
                })
            });
    }

    load(){
        if (this.state.searchQuery !== "") {
            this.searchItems()
        } else {
            this.fetchItems()
        }
    }

    

    render() {
        return (
            <div>
                <Header />  
                <div className="actionsBar">                     
                    <div className="searchBar">                    
                        <Input
                            type="text"
                            name="search"
                            onChange={(e) => {
                                this.state.searchQuery = e.target.value
                                this.setState(this.state)
                                this.load()
                            } }
                            placeholder="Search"                                   
                            value={this.state.searchQuery}                                                          
                        />   
                    </div>                                
                    <a onClick={() => {
                                    this.props.popups.open(GPGModal, {
                                                title: <b>Edit Columns</b>,
                                                content: <Columns 
                                                    columns={this.state.columns} 
                                                    onChange={ (columns) => {
                                                        this.state.columns = [ ...columns]
                                                        const omit = {}
                                                        columns.map((column) => {
                                                            omit[column.identifier] = column.omit
                                                        })
                                                        localStorage.setItem("user_columns", JSON.stringify(omit))
                                                        this.setState(this.state)
                                                    } }
                                                 />,
                                                onClose: (...params) => {
                                                }
                                                }); 
                                }}><MdOutlineViewWeek /> Columns</a>
                    <Link to="/user/edit"><MdCreate /> Create</Link>                                                         
                </div>
                <div className="filters">
                        
                        <div className="clear"></div>
                    </div>
                <div className="tableWrapper" style={{borderTopColor:StringToColour('user')}}>                                                             
                <DataTable
                    title="User"
                    highlightOnHover="true"
                    pointerOnHover="true"
                    pagination="true"
                    paginationServer="true"
                    paginationPerPage={this.state.pageSize}
                    columns={this.state.columns}
                    data={this.state.items}
                    pending={this.state.pending} 
                    paginationTotalRows={this.state.total}                   
			        onChangePage={(page) => {                            
                            this.state.page = page-1 
                            this.setState(this.state)	
                            this.load()
                        }
                    }
                    onChangeRowsPerPage={(size, page) => {
                            this.state.page = page-1
                            this.state.pageSize = size
                            this.setState(this.state)  
                            this.load()                        
                        }
                    }
                    onRowClicked={(row) => {                             
                            if (pressedKeys[91]) {
                                var url = "/user/edit/"+row.id;
                                window.open(url, '_blank', 'noopener,noreferrer').focus();
                            } else {
                                window.location.href = "/user/edit/"+row.id;                           
                            }
                        }
                    }
                />
                </div>
            </div>
        )
    }
}



export default (props) => <User {...props}
            params={useParams()} popups={usePopupManager()} />

class Columns extends React.Component {
    constructor(props) {
        super(props)        
        this.state = {             
            columns: props.columns
        }
    }

    render() {
        return (
            <ul className="columnsEdit">
                {this.state.columns.map((column, index) => 
                    <li key={index} >                                                        
                        <Input
                            type="checkbox"
                            name={column.name}
                            onChange={(e) => {                                
                                this.state.columns[index].omit = !e.target.checked
                                this.setState(this.state)
                                this.props.onChange(this.state.columns)
                            } }                                                                                                    
                            value={ !this.state.columns[index].omit }
                            checked={ !this.state.columns[index].omit }                                                                                                
                        />
                        <span>{column.name}</span>
                    </li>
                )}
            </ul>
        )
    }
}


function EntityRefPromises(items) {
     return items.map((item, index) => { 
        var res = [] 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        return res
    })
}