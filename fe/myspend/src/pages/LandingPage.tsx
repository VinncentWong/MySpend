import { Box } from "@chakra-ui/react";
import FirstContent from "../components/landing/firstcontent";
import Navbar from "../components/navbar";

const LandingPage = () => {
    return(
        <Box>
            <Navbar></Navbar>
            <FirstContent></FirstContent>
        </Box>
    )
};

export default LandingPage;