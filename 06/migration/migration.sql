-- Migration: Create Portfolio Database Schema
-- Description: Creates all necessary tables for the portfolio application

-- Create profile table
CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    tagline VARCHAR(500),
    location VARCHAR(255),
    bio TEXT,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50),
    availability_status VARCHAR(50) DEFAULT 'AVAILABLE',
    github_url VARCHAR(500),
    linkedin_url VARCHAR(500),
    website_url VARCHAR(500),
    resume_url VARCHAR(500),
    profile_image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('LIVE', 'COMPLETED', 'IN_PROGRESS')),
    year INT NOT NULL,
    is_featured BOOLEAN DEFAULT FALSE,
    is_confidential BOOLEAN DEFAULT FALSE,
    project_url VARCHAR(500),
    repository_url VARCHAR(500),
    display_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create technologies table
CREATE TABLE IF NOT EXISTS technologies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    category VARCHAR(100) NOT NULL,
    proficiency_level VARCHAR(50) NOT NULL CHECK (proficiency_level IN ('BEGINNER', 'JUNIOR', 'INTERMEDIATE', 'ADVANCED', 'EXPERT')),
    is_core_technology BOOLEAN DEFAULT FALSE,
    learning_status VARCHAR(50),
    icon_url VARCHAR(500),
    display_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create project_technologies junction table
CREATE TABLE IF NOT EXISTS project_technologies (
    project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    technology_id INT NOT NULL REFERENCES technologies(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (project_id, technology_id)
);

-- Create education table
CREATE TABLE IF NOT EXISTS education (
    id SERIAL PRIMARY KEY,
    institution_name VARCHAR(255) NOT NULL,
    degree_type VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL,
    start_year INT,
    end_year INT,
    field_of_study VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create performance_metrics table
CREATE TABLE IF NOT EXISTS performance_metrics (
    id SERIAL PRIMARY KEY,
    metric_type VARCHAR(100) NOT NULL UNIQUE,
    value VARCHAR(50) NOT NULL,
    label VARCHAR(255) NOT NULL,
    sublabel VARCHAR(255),
    icon_name VARCHAR(100) NOT NULL,
    display_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX idx_projects_status ON projects(status);
CREATE INDEX idx_projects_featured ON projects(is_featured);
CREATE INDEX idx_projects_year ON projects(year);
CREATE INDEX idx_technologies_category ON technologies(category);
CREATE INDEX idx_technologies_core ON technologies(is_core_technology);

-- Create trigger function for updating updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers for auto-updating updated_at
CREATE TRIGGER update_profile_updated_at
    BEFORE UPDATE ON profile
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_projects_updated_at
    BEFORE UPDATE ON projects
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_technologies_updated_at
    BEFORE UPDATE ON technologies
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_education_updated_at
    BEFORE UPDATE ON education
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_performance_metrics_updated_at
    BEFORE UPDATE ON performance_metrics
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Insert sample data for profile (optional)
INSERT INTO profile (name, username, tagline, location, email, phone, availability_status, profile_image_url)
VALUES (
    'Ilmyrat',
    'Gummyyew',
    'Junior Developer',
    'Ashgabat, Turkmenistan',
    'ilmyratgummyyew@gmail.com',
    '+993 (61) 02-51-52',
    'AVAILABLE FOR HIRE',
    '/wwwroot/assets/profile.jpg'
) ON CONFLICT (username) DO NOTHING;

-- Insert sample technologies
INSERT INTO technologies (name, category, proficiency_level, is_core_technology, display_order) VALUES
('WPF', 'Desktop Development', 'JUNIOR', TRUE, 1),
('C#', 'Desktop Development', 'JUNIOR', TRUE, 2),
('Windows Forms', 'Desktop Development', 'JUNIOR', TRUE, 3),
('Golang', 'Web & Backend Development', 'JUNIOR', TRUE, 4),
('HTML & CSS', 'Web & Backend Development', 'BEGINNER', FALSE, 5),
('ASP.NET MVC', 'Web & Backend Development', 'BEGINNER', FALSE, 6),
('MySQL', 'Database Management', 'JUNIOR', TRUE, 7)
('PostgreSQL', 'Database Management', 'JUNIOR', TRUE, 8)
('MongoDB', 'Database Management', 'JUNIOR', TRUE, 9)
ON CONFLICT (name) DO NOTHING;

-- Insert sample performance metrics
INSERT INTO performance_metrics (metric_type, value, label, sublabel, icon_name, display_order) VALUES
('projects_completed', '15', 'Projects Completed', 'WPF & Desktop Apps', 'rocket', 1),
('happy_clients', '11', 'Happy Clients', '4 projects pending', 'users', 2),
('years_experience', '2+', 'Years Experience', 'WPF & C# Development', 'clock', 3),
('technologies', '6', 'Technologies', 'Learning Golang', 'laptop', 4)
ON CONFLICT (metric_type) DO NOTHING;