

-- tlv_hex packet messages table
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    frame_id UUID REFERENCES meter_frames(id) ON DELETE CASCADE,

    total INT,
    flow INT,
    battery INT,
    pressure INT,
    temperature INT,

    magnetic_tamper INT,
    rssi_raw INT,
    serial TEXT,
    valve INT,
    firmware INT,
    network_status INT,

    rtc JSONB,
    extended_status_1a JSONB,
    model TEXT,

    meter_index_20 JSONB,
    counters JSONB,
    ext_block_12 JSONB,
    timestamp_1f JSONB,

    created_at TIMESTAMP DEFAULT NOW()
);


--  meter_frame table

CREATE TABLE meter_frames (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    start_flag TEXT,
    frame_length INT,
    product_type INT,
    meter_address TEXT,
    manufacturer_code TEXT,
    imei TEXT,
    protocol_version INT,
    mid INT,
    encryption_flag INT,
    function_code INT,
    tlv_length INT,
    tlv_hex TEXT,
    checksum TEXT,
    end_flag TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

