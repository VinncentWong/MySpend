import { Box } from "@chakra-ui/react";
import FirstContent from "../components/landing/firstcontent";
import SecondContent from "../components/landing/secondcontent";
import Navbar from "../components/navbar";

const LandingPage = () => {
    return(
        <Box>
            <Navbar></Navbar>
            <FirstContent></FirstContent>
            <SecondContent></SecondContent>
        </Box>
    )
};

export default LandingPage;