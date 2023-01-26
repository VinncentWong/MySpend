import { Box } from "@chakra-ui/react";
import Footer from "../components/footer";
import FirstContent from "../components/landing/firstcontent";
import SecondContent from "../components/landing/secondcontent";
import ThirdContent from "../components/landing/thirdcontent";
import Navbar from "../components/navbar";

const LandingPage = () => {
    return(
        <Box>
            <Navbar></Navbar>
            <FirstContent></FirstContent>
            <SecondContent></SecondContent>
            <ThirdContent></ThirdContent>
            <Footer></Footer>
        </Box>
    )
};

export default LandingPage;