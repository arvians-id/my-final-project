<<<<<<< HEAD
import React from 'react'
import { Box, Flex } from '@chakra-ui/react'
import Navbar from '../components/Navbar'
import Sidebar from '../components/Sidebar'
export default function Home() {
  return (
    <div>
      <Navbar />

      <Flex direction="row" justifyContent="center" alignItems="center">
        <Flex width="20%" minHeight="90vh" bgColor="grey.100" boxShadow='dark-lg'>
          <Sidebar />
        </Flex>

        <Flex width="80%" minHeight="90vh" bg="white">
          <Box m={5}>
            <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
              Home
            </Box>
          </Box>
        </Flex>
      </Flex>
    </div>
  )
=======
import {
    Drawer,
    DrawerBody,
    DrawerHeader,
    DrawerOverlay,
    DrawerContent,
    Button,
    useDisclosure,
    Link,
    Heading,
} from '@chakra-ui/react';
import { HamburgerIcon } from '@chakra-ui/icons';
import React from 'react';
import useStore from '../provider/zustand/store';

export default function Home() {
    const user = useStore((state) => state.user);
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [placement, setPlacement] = React.useState('left');

    return (
        <>
            <Button colorScheme="blue" onClick={onOpen}>
                <HamburgerIcon />
            </Button>
            <Heading>Hy {user.username}, Selamat Datang</Heading>
            <Drawer placement={placement} onClose={onClose} isOpen={isOpen}>
                <DrawerOverlay />
                <DrawerContent>
                    <DrawerHeader borderBottomWidth="1px">
                        Dashboard
                    </DrawerHeader>
                    <DrawerBody>
                        <p>
                            <Link>Series</Link>
                        </p>
                        <p>
                            <Link>Courses</Link>
                        </p>
                        <p>
                            <Link>Account</Link>
                        </p>
                    </DrawerBody>
                </DrawerContent>
            </Drawer>
        </>
    );
>>>>>>> 6a615741ffc3f58b296426e7e401eff7f49bbfa4
}
