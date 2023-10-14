import { createContext, useContext, useState } from "react";
import PropTypes from "prop-types";
import ModalComponent from "../common/ModalComponent";

const ModalContext = createContext();

export function ModalProvider({ children }) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [modalContent, setModalContent] = useState(null);

  const openModal = (content) => {
    setModalContent(content);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setModalContent(null);
  };

  return (
    <ModalContext.Provider value={{ openModal, closeModal }}>
      {children}
      <ModalComponent
        content={modalContent}
        open={isModalOpen}
        handleClose={closeModal}
      />
    </ModalContext.Provider>
  );
}

export function useModal() {
  return useContext(ModalContext);
}

ModalProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
