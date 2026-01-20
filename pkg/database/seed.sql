-- ═══════════════════════════════════════════════════════════════════
-- SEED DATA FOR TESTING
-- User ID: a7d64e79-d36d-4f3d-9e3a-9ae27adecef1
-- ═══════════════════════════════════════════════════════════════════

-- 1. Countries
INSERT INTO countries (id, alpha2, alpha3, name, slug, created_at, updated_at) VALUES
(1, 'EG', 'EGY', 'Egypt', 'egypt', NOW(), NOW()),
(2, 'GB', 'GBR', 'United Kingdom', 'united-kingdom', NOW(), NOW()),
(3, 'BR', 'BRA', 'Brazil', 'brazil', NOW(), NOW());

-- 2. Team Colors ("primary" is a reserved keyword, must be quoted)
INSERT INTO team_colors (id, "primary", secondary, text_color, created_at, updated_at) VALUES
(1, '#C8102E', '#FFFFFF', '#FFFFFF', NOW(), NOW()),
(2, '#FFFFFF', '#C8102E', '#000000', NOW(), NOW());

-- 3. Teams
INSERT INTO teams (id, name, slug, short_name, gender, name_code, country_id, team_colors_id, created_at, updated_at) VALUES
(1, 'Al Ahly SC', 'al-ahly-sc', 'Ahly', 'M', 'AHL', 1, 1, NOW(), NOW()),
(2, 'Zamalek SC', 'zamalek-sc', 'Zamalek', 'M', 'ZAM', 1, 2, NOW(), NOW());

-- 4. Players (15 total: 2G, 5D, 5M, 3F)
-- Goalkeepers (2)
INSERT INTO "Player" (id, "teamId", name, slug, short_name, position, "jerseyNumber", height, gender, country, "dateOfBirthTimestamp", "proposedMarketValueRaw", "shirtNumber", created_at, updated_at) VALUES
(1,  1, 'Mohamed El Shenawy', 'mohamed-el-shenawy', 'El Shenawy', 'G', '1', 191, 'M', 1, '1988-12-18', 2000000, 1, NOW(), NOW()),
(2,  1, 'Ali Lotfi', 'ali-lotfi', 'Lotfi', 'G', '22', 188, 'M', 1, '1995-03-15', 500000, 22, NOW(), NOW()),
(3,  1, 'Ali Maaloul', 'ali-maaloul', 'Maaloul', 'D', '26', 178, 'M', 1, '1990-01-01', 1500000, 26, NOW(), NOW()),
(4,  1, 'Yasser Ibrahim', 'yasser-ibrahim', 'Y. Ibrahim', 'D', '6', 186, 'M', 1, '1993-08-20', 1200000, 6, NOW(), NOW()),
(5,  1, 'Mohamed Abdelmonem', 'mohamed-abdelmonem', 'Abdelmonem', 'D', '5', 190, 'M', 1, '1999-02-10', 3000000, 5, NOW(), NOW()),
(6,  1, 'Akram Tawfik', 'akram-tawfik', 'Tawfik', 'D', '2', 175, 'M', 1, '1997-06-05', 2500000, 2, NOW(), NOW()),
(7,  1, 'Rami Rabia', 'rami-rabia', 'Rabia', 'D', '4', 184, 'M', 1, '1993-04-22', 1000000, 4, NOW(), NOW()),
(8,  1, 'Hamdi Fathi', 'hamdi-fathi', 'Fathi', 'M', '8', 180, 'M', 1, '1998-03-10', 2000000, 8, NOW(), NOW()),
(9,  1, 'Aliou Dieng', 'aliou-dieng', 'Dieng', 'M', '14', 182, 'M', 1, '1997-09-15', 3500000, 14, NOW(), NOW()),
(10, 1, 'Mohamed Magdy Afsha', 'mohamed-magdy-afsha', 'Afsha', 'M', '10', 172, 'M', 1, '1994-11-01', 4000000, 10, NOW(), NOW()),
(11, 1, 'Hussein El Shahat', 'hussein-el-shahat', 'El Shahat', 'M', '7', 170, 'M', 1, '1991-06-20', 1500000, 7, NOW(), NOW()),
(12, 1, 'Emam Ashour', 'emam-ashour', 'Ashour', 'M', '15', 178, 'M', 1, '1998-01-25', 5000000, 15, NOW(), NOW()),
(13, 1, 'Percy Tau', 'percy-tau', 'Tau', 'F', '11', 173, 'M', 2, '1994-05-13', 4500000, 11, NOW(), NOW()),
(14, 1, 'Mohamed Sherif', 'mohamed-sherif', 'Sherif', 'F', '9', 183, 'M', 1, '1996-01-04', 3000000, 9, NOW(), NOW()),
(15, 1, 'Taher Mohamed', 'taher-mohamed', 'Taher', 'F', '17', 176, 'M', 1, '1999-08-12', 2500000, 17, NOW(), NOW());

-- 5. Test User
INSERT INTO users (id, email, username, age, favorite_team_id, favorite_player_id, squad_id, country_id, gender, password_hash, created_at, updated_at) VALUES
('a7d64e79-d36d-4f3d-9e3a-9ae27adecef1', 'testuser@fantasy.com', 'testuser', 25, 1, 13, NULL, 1, 'M', '$2a$10$hashedpassword', NOW(), NOW());
