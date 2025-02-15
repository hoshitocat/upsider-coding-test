CREATE TABLE companies (
    id VARCHAR(255) PRIMARY KEY,
    legal_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    postal_code VARCHAR(8) NOT NULL,
    address TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    company_id VARCHAR(255) NOT NULL REFERENCES companies(id),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE business_partners (
    id VARCHAR(255) PRIMARY KEY,
    company_id VARCHAR(255) NOT NULL REFERENCES companies(id),
    partner_company_id VARCHAR(255) NOT NULL REFERENCES companies(id),
    partner_bank_account_id VARCHAR(255) NOT NULL REFERENCES bank_accounts(id),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (partner_company_id, company_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE bank_accounts (
    id VARCHAR(255) PRIMARY KEY,
    company_id VARCHAR(255) NOT NULL REFERENCES companies(id),
    bank_name VARCHAR(255) NOT NULL,
    branch_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(20) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE invoices (
    id VARCHAR(255) PRIMARY KEY,
    company_id VARCHAR(255) NOT NULL REFERENCES companies(id),
    business_partner_id VARCHAR(255) NOT NULL REFERENCES business_partners(id),
    issue_date DATE NOT NULL,
    payment_amount DECIMAL(10,2) NOT NULL,
    fee_rate DECIMAL(4,2) NOT NULL DEFAULT 4.00,
    fee_amount DECIMAL(10,2) NOT NULL,
    tax_rate DECIMAL(4,2) NOT NULL DEFAULT 10.00,
    tax_amount DECIMAL(10,2) NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    due_date DATE NOT NULL,
    status_id VARCHAR(255) NOT NULL REFERENCES invoice_statuses(id),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_invoices_due_date (due_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE invoice_statuses (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
