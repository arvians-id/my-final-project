import {
    Drawer,
    DrawerBody,
    DrawerHeader,
    DrawerOverlay,
    DrawerContent,
    Button,
    useDisclosure,
    Link,
} from '@chakra-ui/react';
import { HamburgerIcon } from '@chakra-ui/icons';
import React from 'react';

export default function Home() {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [placement, setPlacement] = React.useState('left');

    return (
        <>
            <Button colorScheme="blue" onClick={onOpen}>
                <HamburgerIcon />
            </Button>
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
}
