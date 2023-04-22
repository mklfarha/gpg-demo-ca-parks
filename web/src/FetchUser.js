import axios from 'axios'

const FetchUser = () => {
    var token = localStorage.getItem('token')        
    if (token === undefined || token === null || token === ''){                
        window.location.href = "/login"            
    }

    return axios.post('http://localhost:8080/validate', {}, { headers:{
            'Content-Type': 'application/json',
            'Authorization': 'bearer '+token,
        }
    } )      
}

export default FetchUser
