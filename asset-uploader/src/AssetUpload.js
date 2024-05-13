import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';
import { Upload, Button, Table, Space, Typography, Progress } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
const { Title } = Typography;

function AssetUpload() {
    const [file, setFile] = useState(null);
    const [files, setFiles] = useState([]);
    const [fileList, setFileList] = useState([])
    const [uploading, setUploading] = useState(false);
    const [progress, setProgress] = useState(0);

    const handleFileChange = (file) => {
        setFile(file); // Set the file to state
    };

    const handleUpload = async () => {
        if (!file) {
            alert('Please select a file first.');
            return;
        }
    
        const formData = new FormData();
        formData.append('file', file);
    
        setUploading(true);
        setProgress(0); // Initialize progress at the start of the upload

        try {
            const config = {
                headers: {
                    'Content-Type': 'multipart/form-data'
                },
                onUploadProgress: progressEvent => {
                    // Calculate the percentage of upload completed
                    const percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total);
                    setProgress(percentCompleted); // Update state with the progress
                }
            };

            await axios.post('http://localhost:8080/upload', formData, config);
            fetchFiles(); // Refresh the list after upload
            //setFileList([]); // Clear the file list after upload
            setUploading(false);
        } catch (error) {
            console.error('Error uploading file:', error);
            setUploading(false);
        }
    };

    const handleChange = async (info) => {
        let newFileList = [...info.fileList];
    
        // 1. Limit the number of uploaded files
        // Only to show two recent uploaded files, and old ones will be replaced by the new
        newFileList = newFileList.slice(-2);
    
        // 2. Read from response and show file link
        newFileList = newFileList.map((file) => {
          if (file.response) {
            // Component will show file.url as link
            file.url = file.response.url;
          }
          return file;
        });
    
        setFileList(newFileList);
      };
    

    // Function to fetch the list of files
    const fetchFiles = async () => {
        try {
            const response = await axios.get('http://localhost:8080/files');
            setFiles(response.data);
        } catch (error) {
            console.error('Error fetching files:', error);
        }
    };   
    
    // UseEffect to load files on component mount
    useEffect(() => {
        fetchFiles();
    }, []);  
    
    // Prepare the columns for the Ant Design table
    const columns = [
        {
            title: 'File Name',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Identifier',
            dataIndex: 'stored_filename',
            key: 'stored_filename',
        },
        {
            title: 'Size (bytes)',
            dataIndex: 'size',
            key: 'size',
        },
        {
            title: 'Action',
            key: 'action',
            render: (_, record) => (
                <Space size="middle">
                    <a href={`http://localhost:8080/download/${encodeURIComponent(record.stored_filename)}`}
                       download={record.original_filename + record.file_extension}>
                        Download
                    </a>
                </Space>
            ),
        },
    ];

    // Map files to fit Ant Design's Table data source requirements
    const data = files.map((file, index) => ({
        key: file.id,
        name: file.original_filename + file.file_extension,
        size: file.size,
        stored_filename: file.stored_filename,
        original_filename: file.original_filename,
        file_extension: file.file_extension,
    }));

    return (
        <div style={{ marginTop: '20px', padding: '20px' }}>
            <Title level={2}>Asset Uploader</Title>
            <Space direction="vertical" size="large" style={{ display: 'flex' }}>
            <Upload
                beforeUpload={file => {
                    handleFileChange(file);
                    return false; // Prevent auto uploading
                }}
                onRemove={() => {}}
                fileList={fileList}
                onChange={handleChange}
            >
                <Button icon={<UploadOutlined />} disabled={uploading}>Select File</Button>
            </Upload>
            <Button type="primary" onClick={handleUpload} disabled={fileList.length === 0 || uploading} style={{ marginLeft: 8 }}>
                Upload
            </Button>
            {uploading && <Progress percent={progress} />}
            {files.length > 0 && (
                <Table columns={columns} dataSource={data} />
            )}
            </Space>
        </div>
      );   
}

export default AssetUpload;
