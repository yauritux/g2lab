package com.yauritux.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.yauritux.model.entity.Zone;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 *
 */
public interface ZoneRepository extends JpaRepository<Zone, Long> {

	Zone findByZoneName(String zoneName);
}
