-- ============================================
-- SneaCave — Complete Database Seed
-- Run this ONCE after first docker compose up
-- Order: categories → sneakers → users →
--        carts → cartitems → wishlists →
--        orders → orderitems
-- ============================================


-- ============================================
-- 1. CATEGORIES
-- ============================================
INSERT INTO categories (name, created_at, updated_at) VALUES
('Running',       NOW(), NOW()),
('Basketball',    NOW(), NOW()),
('Lifestyle',     NOW(), NOW()),
('Skateboarding', NOW(), NOW()),
('Training',      NOW(), NOW())
ON CONFLICT DO NOTHING;


-- ============================================
-- 2. SNEAKERS (100 products)
-- category_id: 1=Running 2=Basketball
--              3=Lifestyle 4=Skateboarding 5=Training
-- ============================================
INSERT INTO sneakers (name, brand, category_id, gender, description, price, stock, image_url, is_active, created_at, updated_at) VALUES

-- RUNNING
('Air Zoom Pegasus 40',    'Nike',        1, 'Unisex', 'Versatile daily trainer with responsive cushioning and breathable mesh upper.',                          129.99, 50, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Ultraboost 23',          'Adidas',      1, 'Unisex', 'Energy-returning Boost midsole with Primeknit upper for long-distance comfort.',                         189.99, 40, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Fresh Foam X 1080v12',   'New Balance', 1, 'Unisex', 'Maximum cushioning for long runs with plush Fresh Foam X midsole.',                                      159.99, 35, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Gel-Nimbus 25',          'ASICS',       1, 'Unisex', 'Premium long-distance runner with FF BLAST PLUS ECO cushioning.',                                         159.99, 30, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Clifton 9',              'HOKA',        1, 'Unisex', 'Lightweight maximalist shoe with early-stage Meta-Rocker for smooth transitions.',                         144.99, 45, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Wave Rider 27',          'Mizuno',      1, 'Unisex', 'Smooth and stable ride with ENERZY LITE midsole foam for responsive cushioning.',                          139.99, 25, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Ghost 15',               'Brooks',      1, 'Unisex', 'Soft, smooth, and connected feel underfoot with DNA LOFT v3 cushioning.',                                  139.99, 30, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Kinvara 14',             'Saucony',     1, 'Unisex', 'Lightweight and fast with PWRRUN cushioning for a natural feel.',                                           109.99, 40, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Speedgoat 5',            'HOKA',        1, 'Unisex', 'Trail-ready with Vibram Megagrip outsole and maximum cushioning for mountain running.',                     154.99, 20, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Adrenaline GTS 23',      'Brooks',      1, 'Unisex', 'Supportive stability shoe with GuideRails holistic support system.',                                        139.99, 35, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('React Infinity Run 3',   'Nike',        1, 'Unisex', 'Designed to help reduce injury with a wider base and React foam cushioning.',                               159.99, 30, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Supernova Rise',         'Adidas',      1, 'Unisex', 'Everyday running shoe with DREAMSTRIKE+ cushioning for effortless comfort.',                                119.99, 40, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Kayano 30',              'ASICS',       1, 'Unisex', 'Iconic stability shoe with 4D Guidance System for overpronation support.',                                  179.99, 25, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Endorphin Speed 4',      'Saucony',     1, 'Unisex', 'Carbon-infused nylon plate for speed with PWRRUN PB cushioning.',                                           184.99, 20, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Bondi 8',                'HOKA',        1, 'Unisex', 'Maximum cushioning road shoe with full-EVA midsole for long recovery runs.',                                164.99, 30, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Triumph 21',             'Saucony',     1, 'Unisex', 'Luxuriously cushioned everyday trainer with PWRRUN+ foam.',                                                 149.99, 35, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Glycerin 21',            'Brooks',      1, 'Unisex', 'Plush DNA LOFT v3 cushioning for the ultimate soft landing.',                                               159.99, 28, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Gel-Kayano 30',          'ASICS',       1, 'Male',   'Structured support with Litetruss technology for stability-focused runners.',                               179.99, 22, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Vomero 17',              'Nike',        1, 'Unisex', 'Plush daily trainer with ZoomX foam in the heel for added softness.',                                       169.99, 30, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Solar Glide 6',          'Adidas',      1, 'Unisex', 'Stable everyday trainer with Continental rubber outsole for superior grip.',                                129.99, 40, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),

-- BASKETBALL
('Air Jordan 1 Retro High OG', 'Nike',         2, 'Unisex', 'The iconic silhouette that started it all. Premium leather upper with Air cushioning.',             179.99, 30, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('LeBron 21',                  'Nike',         2, 'Male',   'Built for the king with Max Air cushioning and durable traction pattern.',                           199.99, 25, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Harden Vol. 8',              'Adidas',       2, 'Male',   'Low-cut court shoe with Lightstrike Pro cushioning for quick guards.',                                149.99, 20, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Curry 11',                   'Under Armour', 2, 'Male',   'Engineered for the sharpest shooter in the game with UA Flow cushioning.',                           159.99, 25, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('KD 16',                      'Nike',         2, 'Male',   'Kevin Durant signature with articulated Zoom Air for explosive play.',                               169.99, 20, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Air Jordan 36',              'Nike',         2, 'Male',   'Eclipse Plate technology for multidirectional court responsiveness.',                                 184.99, 18, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Dame 8',                     'Adidas',       2, 'Male',   'Damian Lillard signature with Lightstrike cushioning and herringbone outsole.',                       129.99, 30, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Tatum 2',                    'Jordan',       2, 'Male',   'Jayson Tatum signature with Zoom Air cushioning for two-way players.',                               139.99, 22, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Air Jordan 11 Retro',        'Nike',         2, 'Unisex', 'Patent leather and mesh upper with full-length Air cushioning unit.',                                220.00, 15, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Reebok BB4000 II',           'Reebok',       2, 'Unisex', 'Retro basketball silhouette with leather upper and EVA midsole.',                                    89.99,  35, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Puma MB.03',                 'Puma',         2, 'Male',   'LaMelo Ball signature with Nitro foam for court responsiveness.',                                    129.99, 20, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Air Jordan 4 Retro',         'Nike',         2, 'Unisex', 'Classic 1989 design with visible Air unit and TPU support wings.',                                   209.99, 12, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Giannis Immortality 3',      'Nike',         2, 'Male',   'Accessible performance shoe from the Giannis line with React foam.',                                 79.99,  40, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('D.O.N. Issue 6',             'Adidas',       2, 'Male',   'Donovan Mitchell signature with responsive Bounce cushioning.',                                      109.99, 25, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Air Jordan 3 Retro',         'Nike',         2, 'Unisex', 'First Air Jordan designed by Tinker Hatfield with visible Air heel unit.',                           199.99, 14, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),

-- LIFESTYLE
('Air Force 1 Low',      'Nike',        3, 'Unisex', 'The classic all-white leather sneaker that never goes out of style.',                                        109.99, 80, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Stan Smith',           'Adidas',      3, 'Unisex', 'Minimalist tennis-inspired sneaker with perforated 3-Stripes and leather upper.',                            89.99,  70, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Chuck Taylor All Star','Converse',    3, 'Unisex', 'The original basketball shoe turned cultural icon. Canvas upper, rubber sole.',                              59.99, 100, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Old Skool',            'Vans',        3, 'Unisex', 'Classic skate shoe with the iconic side stripe and durable suede/canvas upper.',                             69.99,  90, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('990v6',                'New Balance', 3, 'Unisex', 'Made in USA premium lifestyle sneaker with ENCAP midsole technology.',                                       184.99, 30, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Gel-Lyte III OG',      'ASICS',       3, 'Unisex', 'Retro runner with split tongue and premium suede upper for a classic look.',                                 109.99, 35, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Suede Classic',        'Puma',        3, 'Unisex', 'Iconic suede upper with Puma Formstrip. A streetwear staple since 1968.',                                    74.99,  60, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Club C 85',            'Reebok',      3, 'Unisex', 'Clean tennis heritage sneaker with leather upper and die-cut EVA midsole.',                                  79.99,  55, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Gazelle',              'Adidas',      3, 'Unisex', 'Vintage-inspired suede sneaker with serrated 3-Stripes and gum rubber sole.',                                99.99,  65, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Air Max 90',           'Nike',        3, 'Unisex', 'Visible Max Air unit in the heel with waffle-pattern outsole. A true classic.',                              129.99, 45, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('990v5',                'New Balance', 3, 'Unisex', 'Premium Made in USA sneaker with pigskin suede and mesh upper.',                                             174.99, 28, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Samba OG',             'Adidas',      3, 'Unisex', 'Indoor football-inspired silhouette with gum sole and T-toe overlay.',                                       99.99,  55, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Air Max 1',            'Nike',        3, 'Unisex', 'The original Air Max with visible cushioning window, designed by Tinker Hatfield.',                          139.99, 40, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Sk8-Hi',               'Vans',        3, 'Unisex', 'High-top canvas sneaker with padded ankle collar and signature side stripe.',                                74.99,  70, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Handball Spezial',     'Adidas',      3, 'Unisex', 'Indoor sport-inspired shoe with suede upper and gum sole. A street favourite.',                              109.99, 35, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Dunk Low',             'Nike',        3, 'Unisex', 'Basketball-turned-lifestyle icon with leather upper and padded collar.',                                     109.99, 50, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Forum Low',            'Adidas',      3, 'Unisex', 'Retro basketball shoe with ankle strap and classic 3-Stripes branding.',                                    99.99,  45, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Air Max 95',           'Nike',        3, 'Unisex', 'Human body-inspired design with layered panels and full-length Air cushioning.',                             174.99, 30, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('574',                  'New Balance', 3, 'Unisex', 'Versatile everyday lifestyle sneaker with ENCAP midsole and suede/mesh upper.',                              84.99,  65, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('One Star OX',          'Converse',    3, 'Unisex', 'Low-profile suede sneaker with single star branding and vulcanized sole.',                                   74.99,  55, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Air Max 97',           'Nike',        3, 'Unisex', 'Full-length Air unit and ripple design inspired by Japanese bullet trains.',                                  174.99, 25, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Campus 00s',           'Adidas',      3, 'Unisex', 'Y2K-inspired low-top with chunky sole and premium suede upper.',                                             109.99, 40, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Classic Leather',      'Reebok',      3, 'Unisex', 'Simple, clean leather sneaker with die-cut EVA midsole for all-day comfort.',                                74.99,  60, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Authentic',            'Vans',        3, 'Unisex', 'The original Vans silhouette. Low-top canvas with vulcanized rubber sole.',                                   59.99,  90, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Air Jordan 1 Low',     'Nike',        3, 'Unisex', 'Low-top version of the iconic Jordan 1 with leather upper and rubber sole.',                                 109.99, 45, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),

-- SKATEBOARDING
('SB Dunk Low Pro',        'Nike',     4, 'Unisex', 'Padded tongue and Zoom Air insole for skate-ready performance and style.',         119.99, 35, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Era',                    'Vans',     4, 'Unisex', 'Double-stitched upper with re-enforced toe cap for skate durability.',             64.99,  70, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Emerica Reynolds G6',    'Emerica',  4, 'Male',   'Slim cupsole construction with G6 foam insole for board feel.',                   74.99,  30, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Busenitz',               'Adidas',   4, 'Male',   'Dennis Busenitz pro model with durable suede and responsive cushioning.',         89.99,  25, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Half Cab',               'Vans',     4, 'Unisex', 'Steve Caballero signature half-cut boot with padded ankle support.',              74.99,  40, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Slip-On Pro',            'Vans',     4, 'Unisex', 'Laceless skate shoe with elastic side accents and waffle outsole.',               64.99,  55, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Jamie Foy Deathwish',    'DC Shoes', 4, 'Male',   'Pro model with IMPACT-I cushioning and abrasion-resistant upper.',                79.99,  20, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Grosso Mid',             'Etnies',   4, 'Male',   'Jeff Grosso tribute shoe with STI foam insole and vulc outsole.',                 69.99,  25, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Coda SC',                'Lakai',    4, 'Male',   'Suede and canvas upper with cushioned footbed for long skate sessions.',          74.99,  22, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Logan',                  'DC Shoes', 4, 'Unisex', 'Low-top skate shoe with pure rubber cup sole and memory foam insole.',            74.99,  30, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),

-- TRAINING
('Metcon 9',               'Nike',         5, 'Unisex', 'Stable and durable cross-training shoe built for weightlifting and HIIT.',       149.99, 40, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Nano X4',                'Reebok',        5, 'Unisex', 'Versatile training shoe with FloatRide Energy Foam and wide toe box.',           139.99, 35, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Adipower Weightlifting 3','Adidas',       5, 'Unisex', 'Olympic lifting shoe with raised heel and stable TPU sole.',                    219.99, 15, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Powerlift 5',            'Adidas',        5, 'Unisex', 'Entry-level weightlifting shoe with raised heel and supportive upper.',          109.99, 30, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Free Metcon 5',          'Nike',          5, 'Unisex', 'Flexible forefoot with stable heel for the ultimate cross-training hybrid.',     129.99, 35, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('F-lite 235 V3',          'Inov-8',        5, 'Unisex', 'Lightweight and flexible gym shoe with POWERFLOW+ midsole.',                    119.99, 20, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Savaleos',               'Under Armour',  5, 'Unisex', 'Weightlifting shoe with wide toe box and rigid heel for maximum stability.',     109.99, 25, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Speed TR 2',             'NOBULL',        5, 'Unisex', 'Durable microfiber trainer with traction lugs for versatile gym performance.',   139.99, 20, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Trainer V2',             'NOBULL',        5, 'Unisex', 'All-surface training shoe with SuperFabric upper and flat rubber outsole.',      129.99, 25, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Minimus TR V1',          'New Balance',   5, 'Unisex', 'Minimal drop training shoe for natural movement during functional fitness.',     99.99,  30, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Lifter 3',               'Nike',          5, 'Unisex', 'Romaleos-inspired lifting shoe with dual strap and elevated heel.',              199.99, 15, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('Nano X3 Adventure',      'Reebok',        5, 'Unisex', 'Trail-ready cross trainer with Lift and Run Chassis for stability.',             149.99, 22, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('Ultrafly 4',             'Under Armour',  5, 'Unisex', 'Trail training shoe with HOVR cushioning and aggressive outsole lugs.',          119.99, 18, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Reebok Legacy Lifter III','Reebok',       5, 'Unisex', 'Competition weightlifting shoe with 22mm heel rise and TPU heel counter.',       199.99, 12, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Conquest TR',            'ASICS',         5, 'Unisex', 'Multi-purpose training shoe with GEL cushioning and durable outsole.',           99.99,  28, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW()),
('Exploud 2',              'Puma',          5, 'Unisex', 'Cross-training shoe with NITRO foam and PWRFrame for stability.',                109.99, 25, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500', true, NOW(), NOW()),
('All Out Blaze Aero TR',  'Merrell',       5, 'Unisex', 'Trail and gym hybrid with Vibram TC5+ outsole and Flexplate technology.',        119.99, 18, 'https://images.unsplash.com/photo-1608231387042-66d1773070a5?w=500', true, NOW(), NOW()),
('UA SlipSpeed',           'Under Armour',  5, 'Unisex', 'Innovative heel-collapse design for seamless gym-to-street transition.',         149.99, 20, 'https://images.unsplash.com/photo-1539185441755-769473a23570?w=500', true, NOW(), NOW()),
('Minimus 40 Trainer',     'New Balance',   5, 'Unisex', 'Low-profile cross trainer with Vibram outsole for multi-directional movement.',  109.99, 22, 'https://images.unsplash.com/photo-1600185365926-3a2ce3cdb9eb?w=500', true, NOW(), NOW()),
('Flexagon Force 4',       'Reebok',        5, 'Unisex', 'Flexible and lightweight training shoe with flex grooves for natural motion.',   79.99,  35, 'https://images.unsplash.com/photo-1606107557195-0e29a4b5b4aa?w=500', true, NOW(), NOW());


-- ============================================
-- 3. USERS (10 sample users)
-- password for all = "password123"
-- bcrypt hash of "password123" (cost 10)
-- ============================================
INSERT INTO users (username, email, password, role, status, is_blocked, created_at, updated_at) VALUES
('john_doe',     'john@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('jane_smith',   'jane@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('mike_jordan',  'mike@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('sarah_connor', 'sarah@example.com',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('alex_turner',  'alex@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('emma_wilson',  'emma@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('chris_brown',  'chris@example.com',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', true,  NOW(), NOW()),
('lisa_park',    'lisa@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('tom_hanks',    'tom@example.com',     '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW()),
('nina_patel',   'nina@example.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lh', 'user',  'active', false, NOW(), NOW())
ON CONFLICT (email) DO NOTHING;


-- ============================================
-- 4. CARTS (one per user, users 2–11)
-- After seeding users, their IDs start from 2
-- because superadmin is already ID=1
-- ============================================
INSERT INTO carts (user_id, created_at, updated_at) VALUES
(2,  NOW(), NOW()),
(3,  NOW(), NOW()),
(4,  NOW(), NOW()),
(5,  NOW(), NOW()),
(6,  NOW(), NOW()),
(7,  NOW(), NOW()),
(8,  NOW(), NOW()),
(9,  NOW(), NOW()),
(10, NOW(), NOW()),
(11, NOW(), NOW())
ON CONFLICT DO NOTHING;


-- ============================================
-- 5. CART ITEMS (sample items in some carts)
-- cart_id matches above, sneaker IDs 1-100
-- ============================================
INSERT INTO cart_items (cart_id, sneaker_id, quantity, created_at, updated_at) VALUES
(1, 1,  2, NOW(), NOW()),
(1, 36, 1, NOW(), NOW()),
(2, 21, 1, NOW(), NOW()),
(2, 5,  1, NOW(), NOW()),
(3, 56, 2, NOW(), NOW()),
(4, 10, 1, NOW(), NOW()),
(4, 71, 1, NOW(), NOW()),
(5, 22, 1, NOW(), NOW()),
(6, 40, 3, NOW(), NOW()),
(7, 15, 1, NOW(), NOW());


-- ============================================
-- 6. WISHLISTS
-- ============================================
INSERT INTO wishlists (user_id, sneaker_id, created_at, updated_at) VALUES
(2,  21, NOW(), NOW()),
(2,  36, NOW(), NOW()),
(2,  56, NOW(), NOW()),
(3,  1,  NOW(), NOW()),
(3,  22, NOW(), NOW()),
(4,  71, NOW(), NOW()),
(4,  85, NOW(), NOW()),
(5,  36, NOW(), NOW()),
(5,  40, NOW(), NOW()),
(6,  21, NOW(), NOW()),
(7,  56, NOW(), NOW()),
(8,  10, NOW(), NOW()),
(8,  15, NOW(), NOW()),
(9,  22, NOW(), NOW()),
(10, 1,  NOW(), NOW()),
(11, 71, NOW(), NOW())
ON CONFLICT DO NOTHING;


-- ============================================
-- 7. ORDERS
-- ============================================
INSERT INTO orders (user_id, total_amount, status, created_at, updated_at) VALUES
(2,  309.98, 'delivered', NOW() - INTERVAL '30 days', NOW()),
(2,  189.99, 'delivered', NOW() - INTERVAL '15 days', NOW()),
(3,  179.99, 'shipped',   NOW() - INTERVAL '5 days',  NOW()),
(4,  299.98, 'paid',      NOW() - INTERVAL '3 days',  NOW()),
(5,  129.99, 'pending',   NOW() - INTERVAL '1 day',   NOW()),
(6,  384.98, 'delivered', NOW() - INTERVAL '45 days', NOW()),
(7,  109.99, 'cancelled', NOW() - INTERVAL '20 days', NOW()),
(8,  219.98, 'shipped',   NOW() - INTERVAL '7 days',  NOW()),
(9,  174.99, 'paid',      NOW() - INTERVAL '2 days',  NOW()),
(10, 139.99, 'pending',   NOW(),                       NOW()),
(11, 359.98, 'delivered', NOW() - INTERVAL '60 days', NOW()),
(2,  109.99, 'delivered', NOW() - INTERVAL '90 days', NOW());


-- ============================================
-- 8. ORDER ITEMS
-- Matches orders above (order IDs 1-12)
-- ============================================
INSERT INTO order_items (order_id, sneaker_id, quantity, price, created_at, updated_at) VALUES
-- Order 1: john_doe - delivered
(1, 1,  1, 129.99, NOW() - INTERVAL '30 days', NOW()),
(1, 36, 1, 109.99, NOW() - INTERVAL '30 days', NOW()),
-- Order 2: john_doe - delivered
(2, 2,  1, 189.99, NOW() - INTERVAL '15 days', NOW()),
-- Order 3: jane_smith - shipped
(3, 21, 1, 179.99, NOW() - INTERVAL '5 days',  NOW()),
-- Order 4: mike_jordan - paid
(4, 5,  1, 144.99, NOW() - INTERVAL '3 days',  NOW()),
(4, 40, 1, 109.99, NOW() - INTERVAL '3 days',  NOW()),
-- Order 5: sarah_connor - pending
(5, 11, 1, 129.99, NOW() - INTERVAL '1 day',   NOW()),
-- Order 6: alex_turner - delivered
(6, 22, 1, 89.99,  NOW() - INTERVAL '45 days', NOW()),
(6, 71, 1, 119.99, NOW() - INTERVAL '45 days', NOW()),
(6, 85, 1, 149.99, NOW() - INTERVAL '45 days', NOW()),
-- Order 7: emma_wilson - cancelled
(7, 44, 1, 109.99, NOW() - INTERVAL '20 days', NOW()),
-- Order 8: chris_brown - shipped
(8, 3,  1, 159.99, NOW() - INTERVAL '7 days',  NOW()),
(8, 23, 1, 59.99,  NOW() - INTERVAL '7 days',  NOW()),
-- Order 9: lisa_park - paid
(9, 50, 1, 174.99, NOW() - INTERVAL '2 days',  NOW()),
-- Order 10: tom_hanks - pending
(10, 82, 1, 139.99, NOW(), NOW()),
-- Order 11: nina_patel - delivered
(11, 21, 1, 179.99, NOW() - INTERVAL '60 days', NOW()),
(11, 56, 1, 109.99, NOW() - INTERVAL '60 days', NOW()),
-- Order 12: john_doe - delivered (old)
(12, 36, 1, 109.99, NOW() - INTERVAL '90 days', NOW());