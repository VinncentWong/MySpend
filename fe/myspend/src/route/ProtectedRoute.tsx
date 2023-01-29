import { Box, Flex } from '@chakra-ui/react';
import { Navigate, Outlet, useNavigate } from 'react-router-dom';
import HomeNavbar from '../components/navbar/home';
import Sidebar from '../components/sidebar';
const ProtectedRoute = () => {
    if(localStorage.getItem("user") == null){
        return <Navigate to="/"/>
    }
    return(
        <Box>
            <HomeNavbar/>
            <Flex
            flexDir='row'
            border='4px solid aqua'>
                <Sidebar/>
                <Outlet/>
            </Flex>
        </Box>
    )
}

export default ProtectedRoute;